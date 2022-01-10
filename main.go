package main

import (
	"flag"
	"fmt"

	"github.com/xiongwei9/face/gen"
	"github.com/xiongwei9/face/util"
)

func main() {
	fmt.Println("hello face")

	input := flag.String("i", "", "input filename")
	outputDir := flag.String("o", "", "output dictionary")
	flag.Parse()

	// 解析thrift文件
	thriftFiles, entryFile, err := gen.ParseThriftFile(*input)
	if err != nil {
		util.Exit(-1, fmt.Sprintf("ParseThriftFile failed: %v", err))
	}

	// 根据thrift生成代码并写入目标目录
	g := gen.NewGenerator(thriftFiles, entryFile, *outputDir)
	err = g.GenCode()
	if err != nil {
		util.Exit(-1, fmt.Sprintf("GenCode failed: %v", err))
	}
	fmt.Println("GenCode success")
}
