{{ $root := . }}
type {{ .Name }} struct {
    {{ range $idx, $field := .Fields }} {{ $field.Name }} {{ $field.Type }}
    {{ end }}
}
