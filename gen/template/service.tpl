{{ $root := . }}
type {{ $root.Name }} struct {}

{{ range $idx, $method := $root.Methods  }}func (s *{{ $root.Name }}) {{ $method.Name }}(c *gin.Context, params *{{ $method.ArgumentType }}) (*{{ $method.ReturnType }}, error) {
    var resp *{{ $method.ReturnType }} = nil // TODO
    return resp, nil
}
{{ end }}

func {{ $root.Name }}_RouteGroup(r *gin.RouterGroup) {
    s := &{{ $root.Name }}{}
{{ range $idx, $method := $root.Methods }}
    r.{{ $method.HttpMethod }}("{{ $method.HttpPath }}", NewMethod(s.{{ $method.Name }})) {{ end }}

}
