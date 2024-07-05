package public

import (
	"github.com/e421083458/gorm"
	"github.com/xiaka53/DeployAndLog/lib"
)

var (
	MainSql *gorm.DB
)

// 数据库初始化
func InitMysql() (err error) {
	if MainSql, err = lib.GetGormPool("base"); err != nil {
		return
	}
	return
}
