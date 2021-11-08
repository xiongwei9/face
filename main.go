package main

import (
	"fmt"
	"os"

	"github.com/samuel/go-thrift/parser"
)

func main() {
	useGoThrift()
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
