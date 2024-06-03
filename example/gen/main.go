package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/example"
	"github.com/lesomnus/entpb/pbgen/ident"
)

// CWD is assumed to be a project root.
func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("get wd: %w", err))
	}

	wd = filepath.Join(wd, "example")

	proto_file_init := entpb.ProtoFileInit{
		PbPackage: ident.Full{"entpb", "directory"},
		GoPackage: "github.com/lesomnus/entpb/pb",
	}

	entpb_ext, err := entpb.NewExtension(filepath.Join(wd, "proto"))
	if err != nil {
		log.Fatal(fmt.Errorf("create entpb extension: %w", err))
	}

	err = entc.Generate(
		filepath.Join(wd, "schema"),
		&gen.Config{
			Package: "gtihub.com/lesomnus/entpb/ent",
			Target:  filepath.Join(wd, "_ent"),
		},
		entc.Extensions(entpb_ext),
		entc.Annotations(
			entpb.ProtoFiles{
				"entpb/directory/common.proto": entpb.NewProtoFile(proto_file_init).
					AddEnum(example.Role(""), []entpb.EnumField{
						{Name: example.RoleOwner, Number: 10},
						{Name: example.RoleMember, Number: 20},
					}),
				"entpb/directory/service.proto": entpb.NewProtoFile(proto_file_init).
					AliasAs("svc"),
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
