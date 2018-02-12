package main

var modelsTmpl = `{{range $key, $model := .}}# {{$key}}

Parameter | Type | Description
--------- | ---- | -----------{{range $y, $field := $model.Fields}}
{{ $field.Name }} | {{ $field.Type }} | {{ $field.Description }}{{end}}

{{end}}
`
