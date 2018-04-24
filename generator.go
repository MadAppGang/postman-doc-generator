package main

import (
	"go/ast"
	"log"
	"path/filepath"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
	"github.com/madappgang/postman-doc-generator/sugar"
)

// Generator represents structures search tool and transformations them to models.
type Generator struct {
	// structNames contains struct names for search
	structNames []string
	// structs will be contains found structs from structNames
	structs map[string]*ast.StructType
}

// NewGenerator creates the new generator.
func NewGenerator(structNames []string) Generator {
	return Generator{
		structs:     make(map[string]*ast.StructType),
		structNames: structNames,
	}
}

// ParseDir method parses .go files from the specified directory.
func (g *Generator) ParseDir(dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		log.Fatalf("Cannot read dir. %v", err)
	}

	files = excludeTestFiles(files)

	g.ParseFiles(files)
}

// ParseFiles method parses specified files by name.
func (g *Generator) ParseFiles(names []string) {
	for _, name := range names {
		g.ParseFile(name)
	}
}

// ParseFile method parses specified file by name and adds necessary structs to the generator.
func (g *Generator) ParseFile(name string) {
	g.structs = ParseFile(name)
}

// GetModels method transformations found structs to models and returns it.
func (g *Generator) GetModels() []models.Model {
	var models []models.Model

	for name, st := range g.structs {
		model := sugar.ParseStruct(name, *st)
		models = append(models, model)
	}

	return models
}

// excludeTestFiles returns a new slice without test files name.
func excludeTestFiles(names []string) []string {
	var filtered []string

	for _, name := range names {
		if !strings.HasSuffix(name, "_test.go") {
			filtered = append(filtered, name)
		}
	}

	return filtered
}
