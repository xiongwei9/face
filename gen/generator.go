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
	err = codeBuilder.BuildPackages(thrift.Namespaces)
	if err != nil {
		return nil, fmt.Errorf("[Generator] BuildPackages failed: %v", err)
	}
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
	outputDir := path.Join(g.OutputDir, c.Package.Path)

	// 判断是否存在目录，若不存在则创建目录
	outputDirInfo, statErr := os.Stat(outputDir)
	if statErr != nil {
		if !os.IsNotExist(statErr) {
			return fmt.Errorf("output dir: %s is exists. but something failed: %v", outputDir, statErr)
		}
		mkdirErr := os.MkdirAll(outputDir, os.ModePerm)
		if mkdirErr != nil {
			return fmt.Errorf("mkdir for %s failed. message: %v", outputDir, mkdirErr)
		}
	} else if !outputDirInfo.IsDir() {
		return fmt.Errorf("%s is not dictionary", outputDir)
	}

	// TODO: path join in Windows
	// 以入口thrift文件名为所生成的go文件名
	goFilename := getGoFilename(g.Entry)
	targetFilename := path.Join(outputDir, goFilename)
	file, fileErr := os.Create(targetFilename)
	if fileErr != nil {
		return fmt.Errorf("create file: %s failed. message: %v", targetFilename, fileErr)
	}

	// 拼接所生成的代码块，写入文件
	codeText := fmt.Sprintf("%s\n%s\n%s\n%s\n", c.Package.Code, c.EnumBuilder.String(), c.StructBuilder.String(), c.ServiceBuilder.String())
	_, writeErr := file.WriteString(codeText)
	if writeErr != nil {
		return fmt.Errorf("write file: %s failed. message: %v", targetFilename, writeErr)
	}
	return nil
}

func getGoFilename(originFilename string) string {
	wholeName := path.Base(originFilename)
	suffix := path.Ext(originFilename)
	return fmt.Sprintf("%s.go", wholeName[0:len(wholeName)-len(suffix)])
}
