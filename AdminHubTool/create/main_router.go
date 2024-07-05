package create

import (
	"fmt"
	"github.com/xiaka53/AdminHub/AdminHubTool/templae"
	"os"
	"text/template"
)

func setMain() {
	tmpl, err := template.New("main_template").Parse(templae.TemplateMainStr)
	if err != nil {
		fmt.Println("解析模板失败,err:", err)
		os.Exit(0)
	}
	fileName := "main.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败,err:", err)
		os.Exit(0)
	}
	defer file.Close()

	// 执行模板并写入文件
	if err = tmpl.Execute(file, nil); err != nil {
		fmt.Println("执行模板失败,err:", err)
		os.Exit(0)
	}

	fmt.Printf("生成的文件：%s\n", fileName)
}

func setRouter() {
	tmpl, err := template.New("router_template").Parse(templae.TemplateRouterStr)
	if err != nil {
		fmt.Println("解析模板失败,err:", err)
		os.Exit(0)
	}
	fileName := "router.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败,err:", err)
		os.Exit(0)
	}
	defer file.Close()

	// 执行模板并写入文件
	if err = tmpl.Execute(file, nil); err != nil {
		fmt.Println("执行模板失败,err:", err)
		os.Exit(0)
	}

	fmt.Printf("生成的文件：%s\n", fileName)
}
