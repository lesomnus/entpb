package cmd

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"entgo.io/ent/schema/field"
	"github.com/go-openapi/inflect"
	"github.com/iancoleman/strcase"
	"github.com/lesomnus/entpb"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	//go:embed *.go.tpl
	template_files embed.FS
)

var (
	importRuntime   protogen.GoImportPath = "github.com/lesomnus/entpb/cmd/protoc-gen-entpb/runtime"
	importStatus    protogen.GoImportPath = "google.golang.org/grpc/status"
	importCodes     protogen.GoImportPath = "google.golang.org/grpc/codes"
	importUuid      protogen.GoImportPath = "github.com/google/uuid"
	importTimestamp protogen.GoImportPath = "google.golang.org/protobuf/types/known/timestamppb"
)

func NewTemplate(g *protogen.GeneratedFile) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"singular": inflect.Singularize,
		"pascal":   strcase.ToCamel,
		"entname": func(name string) string {
			// FIXME: I don't know how Ent make acronym from the arbitrary input
			// but current example only have this so I hard-coded it.
			if name == "id" {
				return "ID"
			}
			return strcase.ToCamel(name)
		},
		"use": g.QualifiedGoIdent,
		"import": func(import_path string) protogen.GoImportPath {
			return protogen.GoImportPath(import_path)
		},
		"ident": func(name string, import_path protogen.GoImportPath) string {
			return g.QualifiedGoIdent(protogen.GoIdent{
				GoName:       name,
				GoImportPath: import_path,
			})
		},
		"to_ent": func(f *entpb.FieldAnnotation, ident_in string, ident_out string, body string) string {
			var t field.Type
			if f.IsEdge() {
				t = field.TypeUUID
			} else {
				t = f.EntInfo.Type
			}
			b := strings.ReplaceAll(body, "@", ident_out)
			switch t {
			case field.TypeUUID:
				r := strings.NewReplacer(
					"$i", ident_in,
					"$o", ident_out,
					"$body", b,
					"$uuid", g.QualifiedGoIdent(importUuid.Ident("FromBytes")),
					"$status", g.QualifiedGoIdent(importStatus.Ident("Errorf")),
					"$code", g.QualifiedGoIdent(importCodes.Ident("InvalidArgument")),
					"$name", string(f.Ident),
				)
				return r.Replace(
					`if $o, err := $uuid($i); err != nil {
						return nil, $status($code, "$name: %s", err)
					} else {
						$body
					}`)
			case field.TypeTime:
				return fmt.Sprintf("%s := %s.AsTime()\n%s", ident_out, ident_in, b)

			case field.TypeEnum:
				o := fmt.Sprintf("toEnt%s(%s)", f.EntInfo.RType.Name, ident_in)
				return strings.ReplaceAll(body, "@", o)
			default:
				return strings.ReplaceAll(body, "@", ident_in)
			}
		},
		"to_pb": func(f *entpb.FieldAnnotation, ident_in string) string {
			t := f.EntInfo.Type
			switch t {
			case field.TypeUUID:
				return fmt.Sprintf("%s[:]", ident_in)
			case field.TypeTime:
				return fmt.Sprintf("%s(%s)", g.QualifiedGoIdent(importTimestamp.Ident("New")), ident_in)
			case field.TypeEnum:
				return fmt.Sprintf("toPb%s(%s)", f.PbType.Ident, ident_in)
			default:
				return ident_in
			}
		},
	})
	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
