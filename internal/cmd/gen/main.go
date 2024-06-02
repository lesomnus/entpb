package main

import (
	"fmt"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/example"
)

// CWD is assumed to be a project root.
func main() {
	proto_file_init := entpb.ProtoFileInit{
		PbPackage: "entpb.directory",
		GoPackage: "github.com/lesomnus/entpb/pb",
	}

	entpb_ext, err := entpb.NewExtension("./proto")
	if err != nil {
		log.Fatal(fmt.Errorf("create entpb extension: %w", err))
	}

	err = entc.Generate(
		"./schema",
		&gen.Config{
			Package: "gtihub.com/lesomnus/entpb/ent",
			Target:  "./ent",
		},
		entc.Extensions(entpb_ext),
		entc.Annotations(
			entpb.ProtoFiles{
				"entpb/directory/common.proto": entpb.NewProtoFile(proto_file_init).
					AddEnum(example.Role(""), example.Role("").Map()),
				"entpb/directory/service.proto": entpb.NewProtoFile(proto_file_init).
					AliasAs("svc"),
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
