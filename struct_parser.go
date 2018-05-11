package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
	"github.com/madappgang/postman-doc-generator/sugar/structtag"
)

// ParseFile parses the file specified by filename
func ParseFile(filename string) map[string]ast.StructType {
	return parse(filename, nil)
}

// ParseSource parses the source code of a single Go source file
func ParseSource(src interface{}) map[string]ast.StructType {
	return parse("", src)
}

// collectStructs inspects specified node, by adding struct types to map and returns it
func collectStructs(node ast.Node) map[string]ast.StructType {
	structs := make(map[string]ast.StructType, 0)

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
		structs[structName] = *st

		return true
	})

	return structs
}

// parse method parses the file specified by filename or the source code of a single Go source file
func parse(filename string, src interface{}) map[string]ast.StructType {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return collectStructs(node)
}

// ParseStruct converts specified struct to model and returns it
func ParseStruct(name string, st ast.StructType) models.Model {
	model := models.NewModel(name)
	var fields []models.Field
	for _, field := range st.Fields.List {
		fieldName := getFieldName(*field)
		fieldType := getExprName(field.Type)
		fieldDesc := getDocsByField(field)
		f := models.NewField(fieldName, fieldType, fieldDesc)
		fields = append(fields, f)
	}
	model.AddField(fields...)
	return model
}

// getFieldName returns name from tag by json key,
// if name is missing returns name of field in lowercase.
func getFieldName(field ast.Field) string {
	if field.Tag == nil {
		return strings.ToLower(field.Names[0].Name)
	}

	tag := strings.Replace(field.Tag.Value, "`", "", -1)
	name := structtag.GetNameFromTag(tag, "json")
	if name == "" {
		name = strings.ToLower(field.Names[0].Name)
	}

	return name
}

// getExprName returns name from the specified expression
func getExprName(expr ast.Expr) string {
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
