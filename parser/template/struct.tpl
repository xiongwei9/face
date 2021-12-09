{{ $root := . }}
type {{ .Name }} struct {
    {{ range $idx, $value := .Values }} {{ $value.Name }} = {{ $value.Value }}
    {{ end }}
}
