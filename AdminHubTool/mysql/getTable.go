package mysql

import (
	"fmt"
	"github.com/e421083458/gorm"
	"github.com/xiaka53/AdminHub/dao"
	"os"
)

func GetTable(sqlConf MysqlConf) []string {
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

	type Table struct {
		TableName string `gorm:"column:TABLE_NAME"`
	}

	var (
		tables  []Table
		_tables []string
	)

	if err = db.Raw("SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = ?", sqlConf.Database).Scan(&tables).Error; err != nil {
		fmt.Println("查询表失败，err：", err)
		os.Exit(0)
	}
	for _, v := range tables {
		if v.TableName == (dao.Setting{}).TableName() ||
			v.TableName == (dao.Admin{}).TableName() ||
			v.TableName == (dao.Menu{}).TableName() ||
			v.TableName == (dao.Role{}).TableName() ||
			v.TableName == (dao.UploadSet{}).TableName() ||
			v.TableName == (dao.UploadFiles{}).TableName() ||
			v.TableName == (dao.AdminLog{}).TableName() {
			continue
		}
		_tables = append(_tables, v.TableName)
	}
	return _tables
}
