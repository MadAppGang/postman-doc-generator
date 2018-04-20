package sugar

import (
	"go/ast"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
)

// ParseStruct converts specified struct to model and returns it
func ParseStruct(name string, st ast.StructType) models.Model {
	model := models.NewModel(name)
	var fields []models.Field
	for _, field := range st.Fields.List {
		fieldName := field.Names[0].Name
		fieldType := getName(field.Type)
		fieldDesc := getDocsByField(field)
		f := models.NewField(fieldName, fieldType, fieldDesc)
		fields = append(fields, f)
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
