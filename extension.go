package entpb

import (
	"fmt"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type Extension struct {
	entc.DefaultExtension
	out_dir string
}

// ExtensionOption is an option for the entproto extension.
type ExtensionOption func(*Extension)

// NewExtension returns a new Extension configured by opts.
func NewExtension(out_dir string, opts ...ExtensionOption) (*Extension, error) {
	if out_dir == "" {
		return nil, fmt.Errorf(`"out_dir" cannot be an empty string`)
	}

	if p, err := filepath.Abs(out_dir); err != nil {
		return nil, fmt.Errorf("resolve absolute path for output directory: %w", err)
	} else {
		out_dir = p
	}

	e := &Extension{out_dir: out_dir}
	for _, opt := range opts {
		opt(e)
	}
	return e, nil
}

// Hooks implements entc.Extension.
func (e *Extension) Hooks() []gen.Hook {
	return []gen.Hook{e.hook()}
}

func (e *Extension) hook() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			// Because Generate has side effects (it is writing to the filesystem under gen.Config.Target),
			// we first run all generators, and only then invoke our code. This isn't great, and there's an
			// [open issue](https://github.com/ent/ent/issues/1311) to support this use-case better.
			if err := next.Generate(g); err != nil {
				return err
			}
			if err := e.generate(g); err != nil {
				return fmt.Errorf("entpb: %w", err)
			}
			return nil
		})
	}
}

func (e *Extension) generate(g *gen.Graph) error {
	if e.out_dir == "" {
		panic(`"out_dir" is empty`)
	}

	gen := generator{
		out_dir: e.out_dir,

		files: map[string]*ProtoFile{},
		enums: map[string]*ProtoFile{},
	}

	return gen.generate(g)
}
