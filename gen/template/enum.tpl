{{ $root := . }}
type {{ $root.Name }} int32

const (
    {{ range $idx, $value := $root.Values }} {{ $root.Name }}{{ $value.Name }} {{ $root.Name }} = {{ $value.Value }}
    {{ end }}
)
