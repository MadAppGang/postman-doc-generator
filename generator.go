package main

import (
	"go/ast"
	"go/build"
	"log"
	"os"
	"path/filepath"

	"github.com/madappgang/postman-doc-generator/models"
	"github.com/madappgang/postman-doc-generator/sugar"
)

// Generator represents structures search tool and transformations them to models.
type Generator struct {
	// structNames contains struct names for search
	structNames []string
	// structs contains found structs from structNames
	structs map[string]ast.StructType
}

// NewGenerator creates the new generator.
func NewGenerator(structNames []string) Generator {
	return Generator{
		structs:     make(map[string]ast.StructType),
		structNames: structNames,
	}
}

// ParseDir method parses go source files from the specified directory.
func (g *Generator) ParseDir(dir string) {
	var names []string
	pkg, err := build.Default.ImportDir(dir, build.IgnoreVendor)
	if err != nil {
		log.Fatalf("cannot process directory %s: %s", dir, err)
	}

	names = append(names, pkg.GoFiles...)
	names = prefixDirectory(dir, names)

	g.ParseFiles(names)
}

// ParseFiles method parses specified files by name.
func (g *Generator) ParseFiles(names []string) {
	for _, name := range names {
		g.ParseFile(name)
	}
}

// ParseFile method parses specified file by name and adds necessary structs to the generator.
func (g *Generator) ParseFile(name string) {
	structs := ParseFile(name)

	if len(g.structNames) > 0 {
		for _, structName := range g.structNames {
			if st, ok := structs[structName]; ok {
				g.structs[structName] = st
			}
		}
	} else {
		for stName, st := range structs {
			g.structs[stName] = st
		}
	}
}

// ParseSource method parses directory or file depending on the source
func (g *Generator) ParseSource(source string) {
	if isDirectory(source) {
		g.ParseDir(source)
	} else {
		g.ParseFile(source)
	}
}

// GetModels method transformations found structs to models and returns it.
func (g *Generator) GetModels() []models.Model {
	var models []models.Model

	for name, st := range g.structs {
		model := sugar.ParseStruct(name, st)
		models = append(models, model)
	}

	return models
}

// isDirectory returns true if the named file is a directory
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// prefixDirectory places the directory name at the beginning of each name in the list.
func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}
