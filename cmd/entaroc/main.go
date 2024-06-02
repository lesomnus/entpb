package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var conf Config
	flag.StringVar(&conf.SchemaPath, "path", "", "path to schema directory")
	flag.Parse()

	if err := run(&conf); err != nil {
		log.Fatalf("entpb: %v", err)
	}

	if *schema_path == "" {
		log.Fatal("entproto: must specify schema path. use entproto -path ./ent/schema")
	}
	abs, err := filepath.Abs(*schema_path)
	if err != nil {
		log.Fatalf("entproto: failed getting absolute path: %v", err)
	}
	graph, err := entc.LoadGraph(*schema_path, &gen.Config{
		Target: filepath.Dir(abs),
	})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.Generate(graph); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}

func run(conf *Config) error {
	if err := conf.Evaluate(); err != nil {
		return fmt.Errorf("evaluate config: %w", err)
	}

	entc.LoadGraph(conf.SchemaPath, &gen.Config{
		Target: ,
	})
}
