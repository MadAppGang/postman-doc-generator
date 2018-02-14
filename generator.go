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
	return &Generator{
		models: make(Models),
	}
}

// ParseAll method parses all *.go files, creates models from structures and returns it
// Non nil verbose error returns if something goes wrong
func (g *Generator) ParseAll(path string) error {
	files, err := filepath.Glob(filepath.Join(path, "*.go"))
	if err != nil {
		return fmt.Errorf("Cannot read dir. %v", err)
	}

	files = g.excludeTestFiles(files)

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
			typeName := node.Name.Name

			g.parseStruct(typeName, node.Type)
		}

		return true
	})

	return nil
}

// parseStruct passes through given node in search of StructType and adds found structure fields to models
func (g *Generator) parseStruct(name string, node ast.Expr) {
	switch subNode := node.(type) {
	case *ast.StructType:
		for _, field := range subNode.Fields.List {
			// try to parse nested structures
			g.parseStruct(name, field.Type)

			// create a field
			createdField := g.createField(field)
			if createdField.Name == "-" {
				continue
			}

			g.models.AddField(name, createdField)
		}
	case *ast.ArrayType:
		// try to find struct in array
		g.parseStruct(name, subNode.Elt)
	}
}

// createField creates a model field by given ast.Field and returns it
func (g *Generator) createField(field *ast.Field) Field {
	var fieldName, fieldType, fieldDescription string

	// try to get parameters from the tag
	if field.Tag != nil {
		tag := field.Tag.Value
		tag = strings.Replace(tag, "`", "", -1) // remove '`' symbols from string
		structTag := reflect.StructTag(tag)

		fieldName = structTag.Get(nameKey)
		if strings.ContainsAny(fieldName, ",") {
			fieldName = getFirstSubstring(fieldName)
		}

		fieldType = structTag.Get(typeKey)
		fieldDescription = structTag.Get(descriptionKey)
	}

	if fieldType != "" {
		fieldType = createMarkdownLink(fieldType)
	}

	return NewField(fieldName, fieldType, fieldDescription)
}

// createMarkdownLink creates markdown link from given string and returns it
func createMarkdownLink(s string) string {
	return fmt.Sprintf("[%s](#%s)", s, strings.ToLower(s))
}

// parseTag parses given tag, creates a tag and returns it
func parseTag(tag string) Field {
	tag = strings.Replace(tag, "`", "", -1) // remove '`' symbols from string
	structTag := reflect.StructTag(tag)

	name := structTag.Get(nameKey)
	if strings.ContainsAny(name, ",") {
		name = getFirstSubstring(name)
	}

	kind := structTag.Get(typeKey) // kind is exported type
	description := structTag.Get(descriptionKey)

	return NewField(name, kind, description)
}

// getSubstring splits given string by comma separator and returns substring by given index
// if the index is greater than the length of the slide, returns empty string
func getSubstring(s string, index int) string {
	slice := strings.Split(s, ",")

	if len(slice) >= index {
		return slice[index]
	}

	return ""
}

// getFirstSubstring returns the first substring with index 0 for given string
func getFirstSubstring(s string) string {
	return getSubstring(s, 0)
}

// GetModels returns created models
func (g *Generator) GetModels() Models {
	return g.models
}

// Insert method replaces given phrase in the file with the models
// Non nil verbose error returns if something goes wrong
func (g *Generator) Insert(filename, phrase string) error {
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

// excludeTestFiles returns a new slice without test files
func (g *Generator) excludeTestFiles(files []string) []string {
	var newSlice []string

	for _, f := range files {
		if !strings.HasSuffix(f, "_test.go") {
			newSlice = append(newSlice, f)
		}
	}

	return newSlice
}
