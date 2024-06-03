package entpb

import (
	"fmt"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lesomnus/entpb/pbgen"
)

type Extension struct {
	entc.DefaultExtension
	out     Fs
	printer Printer
}

// ExtensionOption is an option for the entproto extension.
type ExtensionOption func(*Extension) error

func WithEdition(edition pbgen.Edition) ExtensionOption {
	return func(e *Extension) error {
		p, err := NewPrinter(edition)
		if err != nil {
			return fmt.Errorf("printer: %w", err)
		}

		e.printer = p
		return nil
	}
}

// NewExtension returns a new Extension configured by opts.
func NewExtension(out Fs, opts ...ExtensionOption) (*Extension, error) {
	p, err := NewPrinter(pbgen.Edition2023)
	if err != nil {
		panic("new printer")
	}

	e := &Extension{
		out:     out,
		printer: p,
	}
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
	b, err := NewBuild(g)
	if err != nil {
		return fmt.Errorf("parse graph: %w", err)
	}
	if err := e.printer.Print(b, e.out); err != nil {
		return fmt.Errorf("print: %w", err)
	}

	return nil
}
