package main

import (
	"errors"
	"flag"
	"fmt"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Config struct {
	EntSchemaPath string
	entpb.GrpcPrinterConfig
}

func main() {
	conf := Config{}

	var flags flag.FlagSet
	flags.StringVar(&conf.EntSchemaPath, "schema_path", "", "ent schema path")
	flags.StringVar(&conf.EntPackage, "ent_package", "", "full package name of generate code by Ent")
	flags.StringVar(&conf.ImportPath, "package", "", "full package name of generated code")

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(plugin *protogen.Plugin) error {
		return run(plugin, conf)
	})
}

func run(plugin *protogen.Plugin, conf Config) error {
	plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023
	plugin.SupportedFeatures = uint64(0 |
		pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL |
		pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)

	graph, err := entc.LoadGraph(conf.EntSchemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("load Ent graph: %w", err)
	}

	entpb.ForwardDeclarations(plugin.FilesByPath, graph)

	build, err := entpb.NewBuild(graph)
	if err != nil {
		return fmt.Errorf("parse Ent graph: %w", err)
	}

	errs := []error{}
	for _, file := range plugin.FilesByPath {
		if !file.Generate {
			continue
		}

		for _, service := range file.Services {
			printer := entpb.GrpcPrinter{
				GrpcPrinterConfig: conf.GrpcPrinterConfig,

				Plugin:  plugin,
				File:    file,
				Service: service,
			}
			if err := printer.Print(build); err != nil {
				errs = append(errs, fmt.Errorf(`generate "%s": %w`, service.Desc.Name(), err))
				continue
			}
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
