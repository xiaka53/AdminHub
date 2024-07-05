package table

import (
	"fmt"
	"github.com/xiaka53/AdminHub/AdminHubTool/mysql"
	"github.com/xiaka53/AdminHub/AdminHubTool/templae"
	"os"
	"path/filepath"
	"text/template"
)

func SetDto(tableName string, columns []mysql.Column) {
	var (
		PrimaryKey      string
		PrimaryKeyType  string
		PrimaryKeySmall string
		PrimaryKeyName  string
		Columns         []map[string]string
	)

	for _, column := range columns {
		if column.ColumnKey == "PRI" {
			PrimaryKey = mysql.ToCamelCase(column.ColumnName)
			PrimaryKeyType = mysql.MapColumnType(column.ColumnType)
			PrimaryKeySmall = column.ColumnName
			PrimaryKeyName = column.ColumnComment
		} else {
			Columns = append(Columns, map[string]string{
				"ColumnNameType": mysql.ToCamelCase(column.ColumnName),
				"GoType":         mysql.MapColumnType(column.ColumnType),
				"ColumnName":     column.ColumnName,
			})
		}
	}
	data := struct {
		StructTableName string
		PrimaryKey      string
		PrimaryKeyType  string
		PrimaryKeySmall string
		PrimaryKeyName  string
		Columns         []map[string]string
	}{
		mysql.ToCamelCase(tableName),
		PrimaryKey,
		PrimaryKeyType,
		PrimaryKeySmall,
		PrimaryKeyName,
		Columns,
	}

	tmpl, err := template.New("dto_template").Parse(templae.TemplateDtoStr)
	if err != nil {
		fmt.Println("解析模板失败,err:", err)
		os.Exit(0)
	}
	fileName := fmt.Sprintf("dto/%s.go", tableName)
	if err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		fmt.Println("创建dto文件夹失败,err:", err)
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
