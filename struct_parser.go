package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
)

// StructParser represents a struct parser
type StructParser struct {
	structs map[string]*ast.StructType
	fset    *token.FileSet
}

// NewStructParser creates a new struct parser
func NewStructParser() *StructParser {
	return &StructParser{
		fset: token.NewFileSet(),
	}
}

// ParseFile parses the file specified by filename
func (sp *StructParser) ParseFile(filename string) {
	node, err := parser.ParseFile(sp.fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	sp.structs = collectStructs(node)
}

// ParseSource parses the source code of a single Go source file
func (sp *StructParser) ParseSource(src interface{}) {
	node, err := parser.ParseFile(sp.fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	sp.structs = collectStructs(node)
}

// GetStruct returns the ast.StructType specified by name
// If the struct couldn't be found, the returned AST is nil and the error
func (sp *StructParser) GetStruct(name string) (*ast.StructType, error) {
	if sp.structs[name] == nil {
		return nil, errors.New("not found")
	}

	return sp.structs[name], nil
}

// structToModel converts specified struct to model and returns it
func structToModel(name string, st ast.StructType) *models.Model {
	model := models.NewModel(name)
	var fields []models.Field
	for _, field := range st.Fields.List {
		fieldName := field.Names[0].Name
		fieldType := getName(field.Type)
		fieldDesc := getDocsByField(field)
		f := models.NewField(fieldName, fieldType, fieldDesc)
		fields = append(fields, *f)
	}
	model.AddField(fields...)
	return model
}

// getName returns name from the specified expression
func getName(expr ast.Expr) string {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return ""
	}
	return ident.Name
}

// getDocsByField returns the text of the comment specified by field
func getDocsByField(f *ast.Field) string {
	return strings.TrimSuffix(f.Doc.Text(), "\n")
}

// collectStructs inspects specified node, by adding struct types to map and returns it
func collectStructs(node ast.Node) map[string]*ast.StructType {
	structs := make(map[string]*ast.StructType, 0)

	ast.Inspect(node, func(x ast.Node) bool {
		ts, ok := x.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structName := ts.Name.Name
		structs[structName] = st

		return true
	})

	return structs
}
