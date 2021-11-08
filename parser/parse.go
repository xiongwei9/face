package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/samuel/go-thrift/parser"
)

func ParseThrift(contents string) {
	parser := &parser.Parser{}
	thrift, err := parser.Parse(strings.NewReader(contents))
	if err != nil {
		fmt.Println("parse error")
		os.Exit(1)
	}
	fmt.Printf("parse success: %v", thrift)
}

func ParseThriftFile(pathStr string) {
	p := &parser.Parser{Filesystem: nil}
	parsedThrift, filename, err := p.ParseFile(pathStr)
	if err != nil {
		fmt.Println("parse error")
		os.Exit(1)
	}

	entryThrift, ok := parsedThrift[filename]
	if !ok {
		fmt.Println("parse error")
		os.Exit(1)
	}

	fmt.Printf("parse success: %s %v", filename, entryThrift)
}
