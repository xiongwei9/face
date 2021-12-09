package template

import "github.com/samuel/go-thrift/parser"

type Enum struct {
	Name   string
	Values []parser.EnumValue
}
