package parser

import (
	"fmt"

	"github.com/samuel/go-thrift/parser"
)

type Generator struct {
	thriftFiles map[string]*parser.Thrift
	entry       string
}

func NewGenerator(thriftFiles map[string]*parser.Thrift, entry string) *Generator {
	return &Generator{
		thriftFiles: thriftFiles,
		entry:       entry,
	}
}

func (g *Generator) genThrift() error {
	// thrift := g.thriftFiles[g.entry]
	// for srvName, srv := range thrift.Services {
	// 	fmt.Printf("%s, %v\n", srvName, srv)
	// 	g.genService(srv)
	// }
	// return nil

	entryFile := g.entry
	thrift := g.thriftFiles[entryFile]
	err := g.genThriftFile(thrift)

	// for fileName, thrift := range g.thriftFiles {}

	return err
}

func (g *Generator) genThriftFile(thrift *parser.Thrift) error {
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
