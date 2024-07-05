package create

import (
	"fmt"
	"github.com/xiaka53/AdminHub/AdminHubTool/base"
	"github.com/xiaka53/AdminHub/AdminHubTool/mysql"
	"github.com/xiaka53/AdminHub/AdminHubTool/redis"
	"os"
)

func Create() {
	var (
		adminName = "AdminHub"
		addr      = 8552
		_mysql    mysql.MysqlConf
		_redis    redis.RedisConf
	)

	fmt.Println("请输入项目名称（默认AdminHub）:")
	_, _ = fmt.Scanln(&adminName)
	fmt.Println("请输入项目端口（默认8552）:")
	_, _ = fmt.Scanln(&addr)

	fmt.Println("设置redis:")
	fmt.Println("redis链接（默认127.0.0.1）:")
	_, _ = fmt.Scanln(&_redis.Url)
	fmt.Println("redis端口（默认6379）:")
	_, _ = fmt.Scanln(&_redis.Port)
	fmt.Println("redis密码（默认为空）:")
	_, _ = fmt.Scanln(&_redis.Pass)
	redis.RedisInit(_redis)

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
	mysql.MysqlInit(_mysql, adminName)

	base.WriteToml(adminName, addr)
	setMain()
	setRouter()
	os.Exit(0)
}
