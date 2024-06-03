package entpb

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen"
	"github.com/lesomnus/entpb/pbgen/ident"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/maps"
)

type generator struct {
	out_dir string // Expected it to be a absolute path.

	files map[string]*ProtoFile // Filepath or alias to proto file.

	// Holds proto file that contains the enum definition.
	// Key is global name of the Go type that bound to enum.
	// e.g. for enum "Role" that bound to Go type "Role" in package "github.com/lesomnus/entpb/example",
	// the key would be its global name, "github.com/lesomnus/entpb/example:example.Role".
	// Global name can be build using following functions:
	//   - globalTypeName
	//   - globalTypeNameFromReflect
	//   - globalTypeNameFromEntTypeInfo
	enum_holders map[string]*ProtoFile

	// Holds message definitions.
	// Key is name of Ent Schema.
	// e.g. User, Identity, ...
	schema_to_messages map[string]*messageAnnotation
}

func (g *generator) generate(graph *gen.Graph) error {
	d, ok := decodeAnnotation(&ProtoFiles{}, graph.Annotations)
	if !ok {
		return nil
	}

	if err := g.parseGraph(graph); err != nil {
		return fmt.Errorf("parse graph: %w", err)
	}

	for p := range *d {
		// `g` contains aliased files but `d` doesn't.
		// Iterating `d` means iterating files with no aliased file.
		f, ok := g.files[p]
		if !ok {
			panic("invalid generator state: file not found")
		}

		proto_file := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Package: f.pbPackage,
			Options: []pbgen.Option{pbgen.FeatureFieldPresenceImplicit},
		}
		if f.goPackage != "" {
			proto_file.Options = append(proto_file.Options, pbgen.OptionGoPackage(f.goPackage))
		}

		// Collect imports.
		{
			imports := map[string]struct{}{}
			for _, message := range f.messages {
				for _, field := range message.fields {
					imports[field.pb_type.Import] = struct{}{}
				}
			}
			for _, service := range f.services {
				for _, rpc := range service.Rpcs {
					imports[rpc.Req.Import] = struct{}{}
					imports[rpc.Res.Import] = struct{}{}
				}
			}
			delete(imports, "")
			delete(imports, f.path)

			paths := maps.Keys(imports)
			slices.Sort(paths)
			for _, p := range paths {
				proto_file.Imports = append(proto_file.Imports, pbgen.Import{Name: p})
			}
		}

		//
		// Service definitions.
		//
		services := maps.Values(f.services)
		slices.SortFunc(services, func(a, b *service) int {
			return cmp.Compare(a.Name, b.Name)
		})
		for _, service := range services {
			d := pbgen.Service{Name: service.Name}

			rpcs := maps.Values(service.Rpcs)
			slices.SortFunc(rpcs, func(a, b *Rpc) int {
				return cmp.Compare(a.Name, b.Name)
			})
			for _, v := range rpcs {
				if v.Comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.Comment})
				}

				elem := pbgen.Rpc{Name: v.Name}
				elem.Request.Type = v.Req.ReferencedIn(f.pbPackage)
				elem.Response.Type = v.Res.ReferencedIn(f.pbPackage)
				switch v.Stream {
				case StreamNone:
				case StreamClient:
					elem.Request.Stream = true
				case StreamServer:
					elem.Response.Stream = true
				case StreamBoth:
					elem.Request.Stream = true
					elem.Response.Stream = true

				default:
					// ignore
				}

				d.Body = append(d.Body, elem)
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		//
		// Enum definitions
		//
		enums := maps.Values(f.enums)
		slices.SortFunc(enums, func(a, b *enum) int {
			return cmp.Compare(a.t.Name(), b.t.Name())
		})
		for _, enum := range enums {
			// TODO: comment for enum

			//
			// Field definitions
			//
			d := pbgen.Enum{Name: enum.t.Name()}
			if enum.is_closed {
				d.Options = append(d.Options, pbgen.FeatureEnumTypeClosed)
			}
			fields := slices.Clone(enum.vs)
			slices.SortFunc(fields, func(a, b EnumField) int {
				return cmp.Compare(a.Number, b.Number)
			})
			for _, v := range fields {
				if v.Comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.Comment})
				}
				d.Body = append(d.Body, pbgen.EnumField{Name: v.Name, Number: v.Number})
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		//
		// Message definitions
		//
		messages := maps.Values(f.messages)
		slices.SortFunc(messages, func(a, b *messageAnnotation) int {
			return cmp.Compare(a.name, b.name)
		})
		for _, message := range messages {
			if message.comment != "" {
				proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, pbgen.Comment{Value: message.comment})
			}

			//
			// Field definitions
			//
			d := pbgen.Message{Name: message.name}
			fields := slices.Clone(message.fields)
			slices.SortFunc(fields, func(a, b *fieldAnnotation) int {
				return cmp.Compare(a.Number, b.Number)
			})
			for _, v := range fields {
				if v.comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.comment})
				}

				elem := pbgen.MessageField{
					Type:   v.pb_type.ReferencedIn(f.pbPackage),
					Name:   v.name,
					Number: v.Number,
				}
				if v.isRepeated {
					elem.Labels = append(elem.Labels, pbgen.LabelRepeated)
				} else if v.isOptional {
					// Presence of "repeated" fields are not tracked.
					elem.Options = append(elem.Options, pbgen.FeatureFieldPresenceExplicit)
				}

				d.Body = append(d.Body, elem)
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		os_path := filepath.Join(g.out_dir, p)
		if err := os.MkdirAll(filepath.Dir(os_path), 0755); err != nil {
			return fmt.Errorf(`create directory for proto files: %w`, err)
		}

		w, err := os.Create(os_path)
		if err != nil {
			return fmt.Errorf(`create proto file: %w`, err)
		}
		if err := pbgen.Execute(w, proto_file); err != nil {
			return fmt.Errorf(`generate proto file for "%s": %w`, p, err)
		}
	}

	return nil
}

func (g *generator) parseGraph(graph *gen.Graph) error {
	var d map[string]*ProtoFile
	if a, ok := graph.Annotations[ProtoFilesAnnotation]; !ok {
		return nil
	} else if err := mapstructure.Decode(a, &d); err != nil {
		panic(fmt.Errorf("decode %s: %w", ProtoFilesAnnotation, err))
	}

	for p, f := range d {
		if l, ok := d[f.alias]; ok {
			if f == l {
				// Note that the map is mutated while it is iterated,
				// so newly added element can be visited.
				continue
			}
			return fmt.Errorf(`duplicated alias "%s" for "%s"`, f.alias, p)
		}

		for name := range f.enums {
			if _, ok := g.enum_holders[name]; ok {
				return fmt.Errorf(`multiple definition of enum for same Go type "%s"`, name)
			}

			g.enum_holders[name] = f
		}

		f.path = p
		g.files[p] = f
		g.files[f.alias] = f
	}

	errs := []error{}
	for _, s := range graph.Schemas {
		// Note that `parseMessage` does not parse their fields but only its name
		// since there may be dependencies between messages.
		if err := g.parseMessage(s); err != nil {
			errs = append(errs, fmt.Errorf(`schema "%s": %w`, s.Name, err))
			continue
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse messages declarations: %w", errors.Join(errs...))
	}

	errs = []error{}
	for _, f := range g.enum_holders {
		for _, enum := range f.enums {
			if err := g.normalizeEnum(enum); err != nil {
				errs = append(errs, fmt.Errorf(`normalize enum%s: %w`, enum.t.Name(), err))
			}
		}
	}
	for _, msg := range g.schema_to_messages {
		errs_ := []error{}
		if err := g.parseFields(msg); err != nil {
			errs_ = append(errs_, fmt.Errorf(`parse fields: %w`, err))
		}
		if err := g.parseService(msg); err != nil {
			errs_ = append(errs_, fmt.Errorf(`parse service: %w`, err))
		}
		if len(errs_) > 0 {
			errs = append(errs, fmt.Errorf(`schema "%s": %w`, msg.schema.Name, errors.Join(errs_...)))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse message definitions: %w", errors.Join(errs...))
	}

	return nil
}

func (g *generator) parseMessage(r *load.Schema) error {
	d, ok := decodeAnnotation(&messageAnnotation{}, r.Annotations)
	if !ok {
		return nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = ident.Ident(a.Value)
	} else {
		d.name = ident.Ident(r.Name)
	}
	if a, ok := decodeAnnotation(&schema.CommentAnnotation{}, r.Annotations); ok {
		d.comment = a.Text
	}

	f, ok := g.files[d.Filepath]
	if !ok {
		return fmt.Errorf(`message "%s" references non-exists proto file "%s"`, d.name, d.Filepath)
	}

	if _, ok := f.messages[d.name]; ok {
		return fmt.Errorf(`message name "%s" duplicated with proto file "%s"`, d.name, d.Filepath)
	}

	d.file = f
	d.schema = r
	f.messages[d.name] = d
	g.schema_to_messages[r.Name] = d
	return nil
}

func (g *generator) normalizeEnum(enum *enum) error {
	prefix := ""
	has_zero := false
	if enum.prefix == nil {
		// no prefix
	} else if *enum.prefix == "" {
		prefix = fmt.Sprintf("%s_", enum.t.Name())
	} else {
		prefix = fmt.Sprintf("%s_", *enum.prefix)
	}
	prefix = strings.ToUpper(prefix)

	for k, v := range enum.vs {
		if v.Number == 0 {
			has_zero = true
		}

		name := regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllString(v.Name, `${1}_${2}`)
		name = strings.ToUpper(name)
		name = fmt.Sprintf("%s%s", prefix, name)
		enum.vs[k].Name = name
	}
	if !enum.is_closed && !has_zero {
		enum.vs = append(enum.vs, EnumField{
			Name:   fmt.Sprintf("%sUNSPECIFIED", prefix),
			Number: 0,
		})
	}

	return nil
}

func (g *generator) parseFields(m *messageAnnotation) error {
	errs := []error{}
	for _, field := range m.schema.Fields {
		d, err := g.parseEntField(field)
		if err != nil {
			errs = append(errs, fmt.Errorf(`field "%s": %w`, field.Name, err))
			continue
		}
		if d == nil {
			continue
		}

		m.fields = append(m.fields, d)
	}

	edges := slices.Clone(m.schema.Edges)
	for _, edge := range m.schema.Edges {
		if edge.Ref == nil {
			continue
		}
		if edge.Ref.Type == edge.Type {
			edges = append(edges, edge.Ref)
		}
	}
	for _, edge := range edges {
		d, err := g.parseEntEdge(edge)
		if err != nil {
			errs = append(errs, fmt.Errorf(`edge "%s": %w`, edge.Name, err))
			continue
		}
		if d == nil {
			continue
		}

		m.fields = append(m.fields, d)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (g *generator) parseEntField(r *load.Field) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = a.Value
	} else {
		d.name = ident.Ident(r.Name)
	}

	if r.Info.Type == field.TypeEnum {
		name := globalTypeNameFromEntTypeInfo(r.Info)
		f, ok := g.enum_holders[name]
		if !ok {
			return nil, fmt.Errorf("unregistered enum type: %s", name)
		}

		enum, ok := f.enums[name]
		if !ok {
			panic(fmt.Errorf(`invalid state of generator: enum does not exist on file: "%s"`, name))
		}

		d.pb_type = PbType{
			Name:    ident.Ident(enum.t.Name()),
			Package: f.pbPackage,
			Import:  f.path,
		}
	} else if t := pb_types[int(r.Info.Type)]; t.Name == "" {
		return nil, fmt.Errorf("unsupported type: %s", r.Info.Type.String())
	} else {
		d.pb_type = t
	}

	d.comment = r.Comment
	d.isOptional = r.Nillable

	return d, nil
}

func (g *generator) parseEntEdge(r *load.Edge) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = a.Value
	} else {
		d.name = ident.Ident(r.Name)
	}

	message, ok := g.schema_to_messages[r.Type]
	if !ok {
		return nil, fmt.Errorf(`edge "%s" references a schema "%s" that is not a proto message`, r.Name, r.Type)
	}

	d.pb_type = message.pbType()
	d.comment = r.Comment
	d.isOptional = !r.Required
	d.isRepeated = !r.Unique

	return d, nil
}

func (g *generator) parseService(d *messageAnnotation) error {
	s := d.Service
	if s == nil || len(s.Rpcs) == 0 {
		return nil
	}

	s.message = d
	if s.Filepath == "" {
		s.Filepath = d.Filepath
	}

	f, ok := g.files[s.Filepath]
	if !ok {
		return fmt.Errorf(`service "%s" references non-exists proto file "%s"`, d.name, d.Filepath)
	}
	if s.Name == "" {
		s.Name = ident.Ident(fmt.Sprintf("%sService", d.name))
	}
	if _, ok := f.services[s.Name]; ok {
		return fmt.Errorf(`duplicated service "%s"`, s.Name)
	} else {
		f.services[s.Name] = s
	}

	for _, rpc := range s.Rpcs {
		if rpc.Req.Equal(&PbThis) {
			rpc.Req = d.pbType()
		}
		if rpc.Res.Equal(&PbThis) {
			rpc.Res = d.pbType()
		}

		if rpc.Req.Import == "" {
			return fmt.Errorf(`RPC "%s": parameter type must be message`, rpc.Name)
		}
		if rpc.Res.Import == "" {
			return fmt.Errorf(`RPC "%s": return type must be message`, rpc.Name)
		}
	}

	return nil
}
