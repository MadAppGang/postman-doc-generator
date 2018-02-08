package main

var modelsTmpl = `{{range $i, $model := .}}# {{$model.Name}}

Parameter | Type | Description
--------- | ---- | -----------{{range $y, $field := $model.Fields}}
{{ $field.Name }} | {{ $field.Type }} | {{ $field.Description }}{{end}}
{{end}}
`
