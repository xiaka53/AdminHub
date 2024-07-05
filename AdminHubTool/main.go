package main

import (
	"flag"
	"fmt"
	"github.com/xiaka53/AdminHub/AdminHubTool/create"
	table2 "github.com/xiaka53/AdminHub/AdminHubTool/table"
)

var _init, h, table bool

func main() {
	flag.BoolVar(&_init, "init", false, "实例化项目")
	flag.BoolVar(&table, "table", false, "实例化表格")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&h, "help", false, "帮助")

	flag.Parse()

	// 如果检测到 help 参数
	if _init {
		create.Create()
	}
	if table {
		table2.Table()
	}
	help()
}

func help() {
	fmt.Println("========")
	fmt.Println("--init", "创建项目")
	fmt.Println("--table", "创建表格")
}
