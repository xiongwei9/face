package gen

import (
	"strings"

	"github.com/samuel/go-thrift/parser"
)

func ParseThrift(contents string) (map[string]*parser.Thrift, string, error) {
	p := &parser.Parser{}
	thrift, err := p.Parse(strings.NewReader(contents))
	if err != nil {
		return nil, "", err
	}
	entry := "/tmp/tmp.thrift"
	files := map[string]*parser.Thrift{entry: thrift}
	return files, entry, nil
}

func ParseThriftFile(pathStr string) (map[string]*parser.Thrift, string, error) {
	p := &parser.Parser{Filesystem: nil}
	parsedThrift, filename, err := p.ParseFile(pathStr)
	if err != nil {
		return nil, "", err
	}
	return parsedThrift, filename, nil
}
