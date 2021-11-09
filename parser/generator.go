package parser

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/samuel/go-thrift/parser"
)

type Generator struct {
	thriftFiles map[string]*parser.Thrift
	entry       string
	codeBuilder strings.Builder
}

func NewGenerator(thriftFiles map[string]*parser.Thrift, entry string) *Generator {
	return &Generator{
		thriftFiles: thriftFiles,
		entry:       entry,
		codeBuilder: strings.Builder{},
	}
}

// TODO: method.tpl for a whole router-group
func (g *Generator) genService(srv *parser.Service) error {

	tpl, err := template.ParseFiles("./template/method.tpl")
	if err != nil {
		return err
	}

	for methodName, method := range srv.Methods {

		var apiPath, apiMethod string
		for _, annotation := range method.Annotations {
			switch strings.ToLower(annotation.Name) {
			case "api.get":
				apiMethod = http.MethodGet
				apiPath = annotation.Value
			case "api.post":
				apiMethod = http.MethodPost
				apiPath = annotation.Value
			}
		}

		tpl.Execute(&g.codeBuilder, map[string]string{
			"apiPath":   apiPath,
			"apiMethod": apiMethod,
			"apiName":   methodName,
		})
	}

	fmt.Printf("result:\n%s\n", g.codeBuilder.String())
	return nil
}

func (g *Generator) genThrift() error {
	thrift := g.thriftFiles[g.entry]
	for srvName, srv := range thrift.Services {
		fmt.Printf("%s, %v\n", srvName, srv)
		g.genService(srv)
	}
	return nil
}
