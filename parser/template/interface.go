package template

import "github.com/samuel/go-thrift/parser"

type Enum struct {
	Name   string
	Values []parser.EnumValue
}

type StructField struct {
	Name string
	Type string
}

type Struct struct {
	Name   string
	Fields []StructField
}
