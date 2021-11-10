
type {{ .serviceName }} struct {}

{{ range  }}func (s *{{ .serviceName }}) {{ .methodName }}(c *gin.Context) {

} {{ end }}
