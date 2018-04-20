package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// ParseFile parses the file specified by filename
func ParseFile(filename string) map[string]*ast.StructType {
	return parse(filename, nil)
}

// ParseSource parses the source code of a single Go source file
func ParseSource(src interface{}) map[string]*ast.StructType {
	return parse("", src)
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

// parse method parses the file specified by filename or the source code of a single Go source file
func parse(filename string, src interface{}) map[string]*ast.StructType {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return collectStructs(node)
}
