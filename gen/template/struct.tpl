{{ $root := . }}
type {{ .Name }} struct {
    {{ range $idx, $field := .Fields }} {{ $field.Name }} {{ $field.Type }} `json:"{{ $field.JsonKey }}"`
    {{ end }}
}
