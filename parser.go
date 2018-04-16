package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
)

type Parser struct {
	structs map[string]*ast.StructType
	fset    *token.FileSet
}

func NewParser() *Parser {
	return &Parser{
		fset: token.NewFileSet(),
	}
}

func (p *Parser) ParseFile(filename string) {
	node, err := parser.ParseFile(p.fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	p.structs = collectStructs(node)
}

func (p *Parser) ParseSource(src interface{}) {
	node, err := parser.ParseFile(p.fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	p.structs = collectStructs(node)
}

func (p *Parser) GetStruct(name string) *ast.StructType {
	return p.structs[name]
}

func structToModel(name string, st *ast.StructType) *models.Model {
	model := models.NewModel(name)
	var fields []models.Field
	for _, field := range st.Fields.List {
		fieldName := field.Names[0].Name
		fieldType := getName(field.Type)
		fieldDesc := getCommentByField(field)
		f := models.NewField(fieldName, fieldType, fieldDesc)
		fields = append(fields, *f)
	}
	model.AddField(fields...)
	return model
}

func getName(expr ast.Expr) string {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return ""
	}
	return ident.Name
}

func getCommentByField(f *ast.Field) string {
	return strings.TrimSuffix(f.Doc.Text(), "\n")
}

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
