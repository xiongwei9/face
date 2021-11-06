package main

import (
	"fmt"
	"os"

	thriftParser "github.com/cloudwego/thriftgo/parser"
	"github.com/samuel/go-thrift/parser"
)

func main() {
	useThriftGo()
	useGoThrift()
}

func useThriftGo() {
	parserThrift, err := thriftParser.ParseFile("./idl/service.thrift", nil, true)
	if err != nil {
		fmt.Printf("parse error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("parse ok: %v", parserThrift)
}

func useGoThrift() {
	fmt.Println("hello face")
	p := parser.Parser{Filesystem: nil}
	parsedThrift, filename, err := p.ParseFile("./idl/service.thrift")
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
