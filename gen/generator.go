package gen

import (
	"fmt"

	"github.com/samuel/go-thrift/parser"
)

type Generator struct {
	ThriftFiles map[string]*parser.Thrift
	Entry       string
}

func NewGenerator(thriftFiles map[string]*parser.Thrift, entry string) *Generator {
	return &Generator{
		ThriftFiles: thriftFiles,
		Entry:       entry,
	}
}

func (g *Generator) GenCode() error {
	entryFile := g.Entry
	thrift := g.ThriftFiles[entryFile]
	err := g.genCode(thrift)

	return err
}

func (g *Generator) genCode(thrift *parser.Thrift) error {
	codeBuilder := NewCodeBuilder()
	var err error
	err = codeBuilder.BuildEnums(thrift.Enums)
	if err != nil {
		return fmt.Errorf("[Generator] BuildEnums failed: %v", err)
	}
	err = codeBuilder.BuildStruct(thrift.Structs)
	if err != nil {
		return fmt.Errorf("[Generator] BuildStruct failed: %v", err)
	}
	err = codeBuilder.BuildService(thrift.Services)
	if err != nil {
		return fmt.Errorf("[Generator] BuildService failed: %v", err)
	}
	fmt.Printf("%s\n%s\n%s\n", codeBuilder.EnumBuilder.String(), codeBuilder.StructBuilder.String(), codeBuilder.ServiceBuilder.String())
	return nil
}
