package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb"
)

func main() {
	graph, err := entc.LoadGraph("/workspaces/entaro/example/schema", &gen.Config{})
	if err != nil {
		panic(err)
	}

	build, err := entpb.NewBuild(graph)
	if err != nil {
		panic(err)
	}

	_ = build
}
