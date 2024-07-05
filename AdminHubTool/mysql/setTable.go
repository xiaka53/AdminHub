package mysql

import (
	"fmt"
	"github.com/e421083458/gorm"
	"github.com/xiaka53/AdminHub/AdminHubTool/templae"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Column struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT"`
	Extra         string `gorm:"column:EXTRA"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

func SetTable(tableName string, sqlConf MysqlConf) []Column {
	(&sqlConf).confTesting()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlConf.User, sqlConf.Pass, sqlConf.Url, sqlConf.Port, sqlConf.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("mysql链接失败")
		os.Exit(0)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		fmt.Println("mysql链接失败")
		os.Exit(0)
	}

	// 查询表结构
	var columns []Column
	query := `SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY, COLUMN_DEFAULT, EXTRA, COLUMN_COMMENT
	          FROM information_schema.columns 
	          WHERE table_schema = ? AND table_name = ?`
	if err = db.Raw(query, sqlConf.Database, tableName).Scan(&columns).Error; err != nil {
		fmt.Println("查询表结构失败,err:", err)
		os.Exit(0)
	}

	// 定义模板数据
	structName := ToCamelCase(tableName)
	shortName := strings.ToLower(string(tableName[0]))
	columnInfos := make([]map[string]string, 0)
	// 转换数据库字段类型为 Go 类型
	for _, column := range columns {
		goType := MapColumnType(column.ColumnType)
		columnInfos = append(columnInfos, map[string]string{
			"ColumnNameType": ToCamelCase(column.ColumnName),
			"ColumnName":     column.ColumnName,
			"GoType":         goType,
			"ColumnComment":  column.ColumnComment, // 可以从数据库中获取列的注释
		})
	}

	// 准备模板数据
	data := struct {
		StructName string
		ShortName  string
		TableName  string
		Columns    []map[string]string
	}{
		StructName: structName,
		ShortName:  shortName,
		TableName:  tableName,
		Columns:    columnInfos,
	}

	tmpl, err := template.New("dao_template").Parse(templae.TemplateDaoStr)
	if err != nil {
		fmt.Println("解析模板失败,err:", err)
		os.Exit(0)
	}

	fileName := fmt.Sprintf("dao/%s.go", tableName)
	if err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		fmt.Println("创建dao文件夹失败,err:", err)
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
	return columns
}

// 将数据库列类型映射为 Go 类型
func MapColumnType(columnType string) string {
	switch {
	case strings.Contains(columnType, "int"):
		return "int"
	case strings.Contains(columnType, "varchar"), strings.Contains(columnType, "text"):
		return "string"
	case strings.Contains(columnType, "datetime"), strings.Contains(columnType, "timestamp"):
		return "time.Time"
	case strings.Contains(columnType, "decimal"), strings.Contains(columnType, "float"):
		return "float64"
	case strings.Contains(columnType, "tinyint"), strings.Contains(columnType, "smallint"), strings.Contains(columnType, "mediumint"):
		return "int"
	case strings.Contains(columnType, "char"):
		return "string"
	case strings.Contains(columnType, "blob"):
		return "[]byte"
	default:
		return "string" // 默认使用 string 类型
	}
}

// 将下划线分隔命名转换为大写驼峰命名
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}
