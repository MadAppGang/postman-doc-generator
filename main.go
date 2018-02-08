package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strings"
)

const (
	nameKey        = "json"
	typeKey        = "export"
	descriptionKey = "description"
	filename       = "models/user.go"
)

func main() {
	cfg, err := NewConfig(filename)
	if err != nil {
		log.Fatalf("Cannot create a config. %v", err)
	}

	structures := cfg.GetModels()

	structures.Print()
}

// Config represents a model for storing settings
type Config struct {
	filename string
	file     *ast.File
}

// NewConfig initializes new config by given filename and returns it
// Non nil verbose error returns if something goes wrong
func NewConfig(filename string) (*Config, error) {
	// Create the AST by given filename
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse given file. %v", err.Error())
	}

	config := Config{
		filename: filename,
		file:     file,
	}

	return &config, nil
}

// GetModels method inspects nodes in the file, adds found structures to the models array and returns it
func (c *Config) GetModels() Models {
	var models Models

	ast.Inspect(c.file, func(x ast.Node) bool {
		switch node := x.(type) {
		case *ast.TypeSpec:
			models.Add(node.Name.Name)
		case *ast.StructType:
			for _, field := range node.Fields.List {
				tag := field.Tag.Value
				tag = strings.Replace(tag, "`", "", -1)
				structTag := reflect.StructTag(tag)

				json := structTag.Get(nameKey)
				export := structTag.Get(typeKey)
				description := structTag.Get(descriptionKey)

				models.AddField(json, export, description)
			}
		}

		return true
	})

	return models
}
