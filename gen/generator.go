package gen

import (
	"fmt"
	"os"
	"path"

	"github.com/samuel/go-thrift/parser"
)

type Generator struct {
	ThriftFiles map[string]*parser.Thrift
	Entry       string
	OutputDir   string
}

func NewGenerator(thriftFiles map[string]*parser.Thrift, entry string, outputDir string) *Generator {
	return &Generator{
		ThriftFiles: thriftFiles,
		Entry:       entry,
		OutputDir:   outputDir,
	}
}

func (g *Generator) GenCode() error {
	entryFile := g.Entry
	thrift := g.ThriftFiles[entryFile]
	codeBuilder, err := g.genCode(thrift)
	if err != nil {
		return err
	}

	err = g.writeOutputDir(codeBuilder)
	return err
}

func (g *Generator) genCode(thrift *parser.Thrift) (*CodeBuilder, error) {
	codeBuilder := NewCodeBuilder()
	var err error
	err = codeBuilder.BuildEnums(thrift.Enums)
	if err != nil {
		return nil, fmt.Errorf("[Generator] BuildEnums failed: %v", err)
	}
	err = codeBuilder.BuildStruct(thrift.Structs)
	if err != nil {
		return nil, fmt.Errorf("[Generator] BuildStruct failed: %v", err)
	}
	err = codeBuilder.BuildService(thrift.Services)
	if err != nil {
		return nil, fmt.Errorf("[Generator] BuildService failed: %v", err)
	}
	return codeBuilder, nil
}

func (g *Generator) writeOutputDir(c *CodeBuilder) error {
	// 判断是否存在目录
	outputDirInfo, statErr := os.Stat(g.OutputDir)
	if statErr != nil {
		if !os.IsNotExist(statErr) {
			return statErr

		}
		mkdirErr := os.MkdirAll(g.OutputDir, os.ModePerm)
		if mkdirErr != nil {
			return mkdirErr
		}
	} else if !outputDirInfo.IsDir() {
		return fmt.Errorf("%s is not dictionary", g.OutputDir)
	}

	// TODO: fmt error & path join
	goFilename := getGoFilename(g.Entry)
	targetFilename := path.Join(g.OutputDir, goFilename)
	file, fileErr := os.Create(targetFilename)
	if fileErr != nil {
		return fileErr
	}
	codeText := fmt.Sprintf("%s\n%s\n%s\n", c.EnumBuilder.String(), c.StructBuilder.String(), c.ServiceBuilder.String())
	_, writeErr := file.WriteString(codeText)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func getGoFilename(originFilename string) string {
	wholeName := path.Base(originFilename)
	suffix := path.Ext(originFilename)

	goName := fmt.Sprintf("%s.go", wholeName[0:len(wholeName)-len(suffix)])
	return goName
}
