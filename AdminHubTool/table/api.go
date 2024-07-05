package table

import (
	"fmt"
	"github.com/xiaka53/AdminHub/AdminHubTool/mysql"
	"github.com/xiaka53/AdminHub/AdminHubTool/templae"
	"os"
	"path/filepath"
	"text/template"
)

func setApi(serverName, tableName string, columns []mysql.Column) {
	var (
		PrimaryKey string
		Columns    []map[string]string
	)
	for _, column := range columns {
		if column.ColumnKey == "PRI" {
			PrimaryKey = mysql.ToCamelCase(column.ColumnName)
		} else {
			Columns = append(Columns, map[string]string{
				"ColumnNameType":      mysql.ToCamelCase(column.ColumnName),
				"GoType":              mysql.MapColumnType(column.ColumnType),
				"ColumnName":          column.ColumnName,
				"LowerCamelTableName": "_" + tableName,
			})
		}
	}
	data := struct {
		PrimaryKey          string
		ServerName          string
		StructName          string
		TableName           string
		LowerCamelTableName string
		StructTableName     string
		Columns             []map[string]string
	}{
		PrimaryKey,
		serverName,
		tableName,
		tableName,
		"_" + tableName,
		mysql.ToCamelCase(tableName),
		Columns,
	}

	tmpl, err := template.New("api_template").Parse(templae.TemplateApiStr)
	if err != nil {
		fmt.Println("解析模板失败,err:", err)
		os.Exit(0)
	}
	fileName := fmt.Sprintf("api/%s.go", tableName)
	if err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		fmt.Println("创建api文件夹失败,err:", err)
		os.Exit(0)
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败,err:", err)
		os.Exit(0)
	}
	defer file.Close()

	// 执行模板并写入文件
	if err = tmpl.Execute(file, data); err != nil {
		fmt.Println("执行模板失败,err:", err)
		os.Exit(0)
	}

	fmt.Printf("生成的文件：%s\n", fileName)
}
