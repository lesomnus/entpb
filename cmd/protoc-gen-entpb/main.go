package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/cmd/protoc-gen-entpb/cmd"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var (
		schema_path string
		ent_package string
		import_path string
	)

	var flags flag.FlagSet
	flags.StringVar(&schema_path, "schema_path", "", "ent schema path")
	flags.StringVar(&ent_package, "ent_package", "", "full package name of generate code by Ent")
	flags.StringVar(&import_path, "package", "", "full package name of generated code")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		graph, err := entc.LoadGraph(schema_path, &gen.Config{})
		if err != nil {
			return fmt.Errorf("load Ent graph: %w", err)
		}

		entpb.ForwardDeclarations(plugin.FilesByPath, graph)

		build, err := entpb.NewBuild(graph)
		if err != nil {
			return fmt.Errorf("parse Ent graph: %w", err)
		}

		plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
		plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023
		plugin.SupportedFeatures = uint64(0 |
			pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL |
			pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)

		_, name := filepath.Split(import_path)
		p := cmd.Printer{
			EntPackage:  protogen.GoImportPath(ent_package),
			ImportPath:  protogen.GoImportPath(import_path),
			PackageName: name,

			Build:  build,
			Plugin: plugin,
		}
		if err := p.Print(); err != nil {
			return err
		}
		return nil
	})
}
