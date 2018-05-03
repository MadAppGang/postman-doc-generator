package sugar

import (
	"go/ast"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
	"github.com/madappgang/postman-doc-generator/sugar/structtag"
)

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
	name := structtag.GetName(tag, "json")
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
