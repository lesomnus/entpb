package cmd

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"entgo.io/ent/entc/load"
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

func (p *Printer) NewTemplate(g *protogen.GeneratedFile) *template.Template {
	to_ent_with_rv := func(f *entpb.FieldAnnotation, ident_in string, ident_out string, body string, rv string) string {
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
				"$rv", rv,
			)
			return r.Replace(
				`if $o, err := $uuid($i); err != nil {
					return $rv, $status($code, "$name: %s", err)
				} else {
					$body
				}`)
		case field.TypeTime:
			if ident_in[0] == '*' {
				ident_in = ident_in[1:]
			}
			return fmt.Sprintf("%s := %s.AsTime()\n%s", ident_out, ident_in, b)

		case field.TypeEnum:
			o := fmt.Sprintf("toEnt%s(%s)", f.EntInfo.RType.Name, ident_in)
			return strings.ReplaceAll(body, "@", o)

		default:
			return strings.ReplaceAll(body, "@", ident_in)
		}
	}

	t := template.New("")
	t.Funcs(template.FuncMap{
		"singular": inflect.Singularize,
		"plural":   inflect.Pluralize,
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
		"use_ent_type": func(t *field.TypeInfo) string {
			// In fact, all cases are equivalent to:
			// return t.Ident
			switch t.Type {
			case field.TypeTime:
				return g.QualifiedGoIdent(importTimestamp.Ident("Timestamp"))
			case field.TypeUUID:
				return g.QualifiedGoIdent(importUuid.Ident("UUID"))
			default:
				return t.Ident
			}
		},
		"import": func(import_path string) protogen.GoImportPath {
			return protogen.GoImportPath(import_path)
		},
		"ident": func(name string, import_path protogen.GoImportPath) string {
			return g.QualifiedGoIdent(protogen.GoIdent{
				GoName:       name,
				GoImportPath: import_path,
			})
		},
		"schema": func(s *load.Schema) protogen.GoImportPath {
			p := fmt.Sprintf("%s/%s", string(p.EntPackage), strings.ToLower(s.Name))
			return protogen.GoImportPath(p)
		},
		"to_ent_with_rv": to_ent_with_rv,
		"to_ent": func(f *entpb.FieldAnnotation, ident_in string, ident_out string, body string) string {
			return to_ent_with_rv(f, ident_in, ident_out, body, "nil")
		},
		"to_pb": func(f *entpb.FieldAnnotation, ident_in string, ident_out string) string {
			t := f.EntInfo.Type
			switch t {
			case field.TypeUUID:
				return fmt.Sprintf("%s = %s[:]", ident_out, ident_in)

			case field.TypeTime:
				if !f.IsOptional {
					return fmt.Sprintf("%s = %s(%s)",
						ident_out, g.QualifiedGoIdent(importTimestamp.Ident("New")), ident_in)
				}
				return fmt.Sprintf(
					`if %s != nil {
						%s = %s(*%s)
					}`,
					ident_in, ident_out, g.QualifiedGoIdent(importTimestamp.Ident("New")), ident_in)

			case field.TypeEnum:
				return fmt.Sprintf("%s = toPb%s(%s)", ident_out, f.PbType.Ident, ident_in)

			default:
				return fmt.Sprintf("%s = %s", ident_out, ident_in)
			}
		},
		"status": func(code string, msg string) string {
			return fmt.Sprintf(`%s(%s, "%s")`,
				g.QualifiedGoIdent(importStatus.Ident("Errorf")),
				g.QualifiedGoIdent(importCodes.Ident(code)),
				msg,
			)
		},
		"status_err": func(code string, msg string) string {
			return fmt.Sprintf(`%s(%s, "%s: %%s", err)`,
				g.QualifiedGoIdent(importStatus.Ident("Errorf")),
				g.QualifiedGoIdent(importCodes.Ident(code)),
				msg,
			)
		},
	})
	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
