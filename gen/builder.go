package gen

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/ginkgoch/godash"
	"github.com/gobuffalo/packr"
	"github.com/samuel/go-thrift/parser"
	face_template "github.com/xiongwei9/face/gen/template"
)

var PrimitiveTypeThrift2Go = map[string]string{
	"bool":   "bool",
	"byte":   "byte",
	"i32":    "int32",
	"i64":    "int64",
	"double": "float32",
	"string": "string",
}

type PackageInfo struct {
	Code string
	Path string
}

type CodeBuilder struct {
	StructBuilder  strings.Builder
	EnumBuilder    strings.Builder
	ServiceBuilder strings.Builder
	ImportBuilder  strings.Builder
	Package        *PackageInfo
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		StructBuilder:  strings.Builder{},
		EnumBuilder:    strings.Builder{},
		ServiceBuilder: strings.Builder{},
		ImportBuilder:  strings.Builder{},
		Package:        nil,
	}
}

func parseTemplate(filename string) (*template.Template, error) {
	box := packr.NewBox("./template")
	f, err := box.Open(filename)
	if err != nil {
		return nil, err
	}
	// var text []byte
	text, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	t := template.New(filename)
	return t.Parse(string(text))
}

func typeTranslate(thriftType *parser.Type) (string, error) {
	if fieldType, ok := PrimitiveTypeThrift2Go[thriftType.Name]; ok {
		return fieldType, nil
	}

	switch thriftType.Name {
	case "list":
		valType, err := typeTranslate(thriftType.ValueType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("[]%s", valType), nil
	case "set":
		return "", fmt.Errorf("[typeTranslate] NOT support <set>")
	case "map":
		var err error
		keyType, err := typeTranslate(thriftType.KeyType)
		if err != nil {
			return "", err
		}
		valType, err := typeTranslate(thriftType.ValueType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("map[%s]%s", keyType, valType), nil
	}

	return thriftType.Name, nil
	// return "", fmt.Errorf("[typeTranslate]not found any type")
}

func (b *CodeBuilder) BuildService(services map[string]*parser.Service) error {
	tpl, err := parseTemplate("service.tpl")
	if err != nil {
		return err
	}

	for name, service := range services {
		methods := make([]face_template.ServiceMethod, 0, len(service.Methods))

		for methodName, method := range service.Methods {
			var httpPath, httpMethod string
			for _, annotation := range method.Annotations {
				switch strings.ToLower(annotation.Name) {
				case "api.get":
					httpMethod = http.MethodGet
					httpPath = annotation.Value
				case "api.post":
					httpMethod = http.MethodPost
					httpPath = annotation.Value
				}
			}
			returnType, err := typeTranslate(method.ReturnType)
			if err != nil {
				return err
			}
			argumentType, err := typeTranslate(method.Arguments[0].Type)
			if err != nil {
				return err
			}

			serviceMethod := face_template.ServiceMethod{
				Name:         methodName,
				HttpPath:     httpPath,
				HttpMethod:   httpMethod,
				ReturnType:   returnType,
				ArgumentType: argumentType,
			}
			methods = append(methods, serviceMethod)

		}
		serviceData := face_template.Service{
			Name:    name,
			Methods: methods,
		}
		err := tpl.Execute(&b.ServiceBuilder, serviceData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *CodeBuilder) BuildStruct(structs map[string]*parser.Struct) error {
	tpl, err := parseTemplate("struct.tpl")
	if err != nil {
		return err
	}
	for name, structure := range structs {
		fields := make([]face_template.StructField, 0, len(structure.Fields))
		for _, field := range structure.Fields {
			fieldType, err := typeTranslate(field.Type)
			if err != nil {
				return err
			}

			name, nameErr := godash.CamelCaseWithInit(field.Name, true)
			if nameErr != nil {
				name = field.Name
			}
			structField := face_template.StructField{
				Name:    name,
				Type:    fieldType,
				JsonKey: field.Name,
			}
			fields = append(fields, structField)
		}
		structData := &face_template.Struct{
			Name:   name,
			Fields: fields,
		}
		err := tpl.Execute(&b.StructBuilder, structData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *CodeBuilder) BuildEnums(enums map[string]*parser.Enum) error {
	tpl, err := parseTemplate("enum.tpl")
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

func (b *CodeBuilder) BuildPackages(namespaces map[string]string) error {
	for key, pack := range namespaces {
		if key == "go" {
			packagePath := strings.ReplaceAll(pack, ".", "/")
			packageSplit := strings.Split(pack, ".")

			b.Package = &PackageInfo{
				Code: fmt.Sprintf("package %s", packageSplit[len(packageSplit)-1]),
				Path: packagePath,
			}
			return nil
		}
	}
	return nil
}

func (b *CodeBuilder) BuildImportList() error {
	// TODO: ??????NewMethod??????
	tpl, err := parseTemplate("import.tpl")
	if err != nil {
		return err
	}
	err = tpl.Execute(&b.ImportBuilder, nil)
	return err
}
