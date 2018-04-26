package sugar

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/madappgang/postman-doc-generator/models"
)

// ParseStruct converts specified struct to model and returns it
func ParseStruct(name string, st ast.StructType) models.Model {
	model := models.NewModel(name)
	var fields []models.Field
	for _, field := range st.Fields.List {
		if !isExported(field.Names[0].Name) {
			continue
		}
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
	switch x := expr.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getName(x.Elt))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getName(x.Key), getName(x.Value))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s", getName(x.Sel))
	case *ast.StarExpr:
		return getName(x.X)
	case *ast.StructType:
		return "struct"
	}
	return ""
}

// isExported returns true if the first character is capital letter.
func isExported(name string) bool {
	return 'A' <= name[0] && name[0] <= 'Z'
}

// getDocsByField returns the text of the comment specified by field
func getDocsByField(f *ast.Field) string {
	return strings.TrimSuffix(f.Doc.Text(), "\n")
}
