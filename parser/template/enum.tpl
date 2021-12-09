{{ $root := . }}
const (
    {{ range $idx, $value := .Values }} {{ $root.Name }}{{ $value.Name }} = {{ $value.Value }}
    {{ end }}
)
