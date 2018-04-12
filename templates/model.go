package templates

var (
	ModelTmpl = `{{$.Name}}
Parameter | Type | Description
--------- | ---- | -----------{{range $i, $field := $.Fields}}
{{ $field.Name }} | {{ $field.Type }} | {{ $field.Description }}{{end}}`
)
