package parser

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/samuel/go-thrift/parser"
	face_template "github.com/xiongwei9/face/parser/template"
)

type CodeBuilder struct {
	StructBuilder  strings.Builder
	EnumBuilder    strings.Builder
	ServiceBuilder strings.Builder
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		StructBuilder:  strings.Builder{},
		EnumBuilder:    strings.Builder{},
		ServiceBuilder: strings.Builder{},
	}
}

func (b *CodeBuilder) BuildStruct(structs map[string]*parser.Struct) error {

	return nil
}

// TODO: method.tpl for a whole router-group
func (b *CodeBuilder) BuildService(srv *parser.Service) error {

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

		tpl.Execute(&b.ServiceBuilder, map[string]string{
			"apiPath":   apiPath,
			"apiMethod": apiMethod,
			"apiName":   methodName,
		})
	}

	fmt.Printf("result:\n%s\n", b.ServiceBuilder.String())
	return nil
}

func (b *CodeBuilder) BuildEnums(enums map[string]*parser.Enum) error {
	tpl, err := template.ParseFiles("./template/enum.tpl")
	if err != nil {
		return err
	}
	for name, enum := range enums {
		values := make([]parser.EnumValue, 0, len(enum.Values))
		for _, value := range enum.Values {
			values = append(values, *value)
		}
		enumData := &face_template.Enum{
			Name:   name,
			Values: values,
		}
		err := tpl.Execute(&b.EnumBuilder, enumData)
		if err != nil {
			return err
		}
	}
	return nil
}
