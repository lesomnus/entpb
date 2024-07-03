package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/internal/example/role"
	"github.com/lesomnus/entpb/pbgen"
	"github.com/lesomnus/entpb/pbgen/ident"
	"github.com/spf13/afero"
)

var header = []byte(`// Code generated by "github.com/lesomuns/entpb/pbgen". DO NOT EDIT.

`)

type fs_adaptor struct {
	afero.Fs
}

func (a *fs_adaptor) Create(name string) (io.WriteCloser, error) {
	f, err := a.Fs.Create(name)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(header)
	return f, err
}

// CWD is assumed to be a project root.
func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("get wd: %w", err))
	}

	wd = filepath.Join(wd, "internal/example")

	editions := []pbgen.Edition{
		pbgen.Edition2023,
		pbgen.SyntaxProto3,
	}
	for _, edition := range editions {
		base := filepath.Join(wd, "protos", fmt.Sprintf("%s_%s", edition.Keyword, edition.Value))
		fs := afero.NewBasePathFs(afero.NewOsFs(), base)

		entpb_ext, err := entpb.NewExtension(
			&fs_adaptor{Fs: fs},
			entpb.WithEdition(edition),
		)
		if err != nil {
			log.Fatal(fmt.Errorf("create entpb extension: %w", err))
		}

		proto_file_init := entpb.ProtoFileInit{
			PbPackage: ident.Full{"entpb", "directory"},
			GoPackage: "github.com/lesomnus/entpb/internal/example/pb",
		}
		err = entc.Generate(
			filepath.Join(wd, "schema"),
			&gen.Config{
				Package: "github.com/lesomnus/entpb/internal/example/ent",
				Target:  filepath.Join(wd, "ent"),
			},
			entc.Extensions(entpb_ext),
			entc.Annotations(
				entpb.ProtoFiles{
					"khepri/horus/store.proto": entpb.NewProtoFile(proto_file_init),
					"khepri/horus/common.proto": entpb.NewProtoFile(proto_file_init,
						entpb.WithEnum(map[role.Role]entpb.EnumDesc{
							role.Owner:  {Number: role.Owner.V()},
							role.Admin:  {Number: role.Admin.V()},
							role.Member: {Number: role.Member.V()},
						}),
					),
				},
			),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
