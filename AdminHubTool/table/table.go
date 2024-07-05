package table

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/xiaka53/AdminHub/AdminHubTool/mysql"
	"os"
)

func Table() {
	var (
		_mysql     mysql.MysqlConf
		serverName string
		tableName  string
	)
	fmt.Println("项目名称（go module）:")
	_, _ = fmt.Scanln(&serverName)
	fmt.Println("设置mysql:")
	fmt.Println("mysql链接（默认127.0.0.1）:")
	_, _ = fmt.Scanln(&_mysql.Url)
	fmt.Println("mysql端口（默认3306）:")
	_, _ = fmt.Scanln(&_mysql.Port)
	fmt.Println("mysql账号（默认root）:")
	_, _ = fmt.Scanln(&_mysql.User)
SET_MYSQLPASS:
	if _mysql.Pass == "" {
		fmt.Println("mysql密码:")
		_, _ = fmt.Scanln(&_mysql.Pass)
		goto SET_MYSQLPASS
	}
SET_MYSQLDATABASE:
	if _mysql.Database == "" {
		fmt.Println("mysql库:")
		_, _ = fmt.Scanln(&_mysql.Database)
		goto SET_MYSQLDATABASE
	}
	tables := mysql.GetTable(_mysql)
	if len(tables) == 0 {
		fmt.Println("没有可生成的表，请先创建表格")
		os.Exit(0)
	}
	prompt := &survey.Select{
		Message: "请选择要生成的表：",
		Options: tables,
	}
	if err := survey.AskOne(prompt, &tableName); err != nil {
		fmt.Println("请选择表格表格")
		os.Exit(0)
	}
	columns := mysql.SetTable(tableName, _mysql)
	setApi(serverName, tableName, columns)
	SetDto(tableName, columns)
	os.Exit(0)
}
