package parser

import (
	"strings"

	"github.com/samuel/go-thrift/parser"
)

func ParseThrift(contents string) error {
	p := &parser.Parser{}
	thrift, err := p.Parse(strings.NewReader(contents))
	if err != nil {
		return err
	}

	entry := "/tmp/tmp.thrift"
	files := map[string]*parser.Thrift{entry: thrift}
	g := NewGenerator(files, entry)
	g.genThrift()

	return nil
}

func ParseThriftFile(pathStr string) error {
	p := &parser.Parser{Filesystem: nil}
	parsedThrift, filename, err := p.ParseFile(pathStr)
	if err != nil {
		return err
	}

	g := NewGenerator(parsedThrift, filename)
	g.genThrift()

	return nil
}
