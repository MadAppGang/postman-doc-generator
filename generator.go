package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	nameKey        = "json"
	typeKey        = "export"
	descriptionKey = "description"
)

// Generator represents a model for storing parser settings
type Generator struct {
	models Models
}

// NewGenerator initializes new config by given filename and returns it
func NewGenerator() *Generator {
	return &Generator{}
}

// ParseAll method parses all *.go files, creates models from structures and returns it
// Non nil verbose error returns if something goes wrong
func (g *Generator) ParseAll(path string) error {
	files, err := filepath.Glob(filepath.Join(path, "*.go"))
	if err != nil {
		return fmt.Errorf("Cannot read dir. %v", err)
	}

	for _, file := range files {
		isDir, err := isDir(file)
		if err != nil {
			return fmt.Errorf("Cannot get file information. %v", err)
		} else if isDir {
			continue
		}

		err = g.Parse(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// Parse method inspects nodes in the file, adds found structures to the models array and returns nil
// Non nil verbose error returns if something goes wrong
func (g *Generator) Parse(filename string) error {
	// Create the AST by given filename
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return fmt.Errorf("Cannot parse given file. %v", err.Error())
	}

	ast.Inspect(file, func(x ast.Node) bool {
		switch node := x.(type) {
		case *ast.TypeSpec:
			g.models.Add(node.Name.Name)
		case *ast.StructType:
			for _, field := range node.Fields.List {
				if field.Tag == nil {
					// when tag is missing
					continue
				}

				tag := field.Tag.Value
				tag = strings.Replace(tag, "`", "", -1) // remove '`' symbols from string
				structTag := reflect.StructTag(tag)

				json := structTag.Get(nameKey)
				export := structTag.Get(typeKey)
				description := structTag.Get(descriptionKey)

				g.models.AddField(json, export, description)
			}
		}

		return true
	})

	return nil
}

// Inject method replaces given phrase in the file with the models
// Non nil verbose error returns if something goes wrong
func (g *Generator) Inject(filename, phrase string) error {
	models := fmt.Sprintf("%q", g.models)
	models = strings.Replace(models, "\"", "", -1)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read the file. %v", err)
	}

	newContents := strings.Replace(string(contents), phrase, models, -1)

	return ioutil.WriteFile(filename, []byte(newContents), 0)
}

// isDir method gets information about the path and returns true or false
// Non nil verbose error returns if something goes wrong
func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	return fileInfo.IsDir(), err
}
