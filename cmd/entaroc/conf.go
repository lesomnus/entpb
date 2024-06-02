package main

import (
	"fmt"
	"path/filepath"
)

type Config struct {
	SchemaPath string
}

func (c *Config) Evaluate() error {
	if c.SchemaPath == "" {
		return fmt.Errorf("schema path not provided")
	} else if abs, err := filepath.Abs(c.SchemaPath); err != nil {
		return fmt.Errorf("get absolute path to schema")
	} else {
		c.SchemaPath = abs
	}

	return nil
}
