package dao

import (
	"github.com/xiaka53/AdminHub/public"
	"time"
)

type AdminLog struct {
	Id      uint      `gorm:"column:id;primaryKey;type:int(12) auto_increment;comment:'ID'"`
	AdminId uint      `gorm:"column:admin_id;type:int(12);comment:'管理员Id'"`
	Ip      string    `gorm:"column:ip;type:varchar(15);not null;comment:'ip'"`
	Address string    `gorm:"column:address;type:varchar(15);not null;comment:'地址'"`
	Atime   time.Time `gorm:"column:atime;type:datetime;comment:'添加时间'"`
	Route   string    `gorm:"column:route;type:char(32);not null;comment:'路由'"`
	Desc    string    `gorm:"column:desc;type:char(64);not null;comment:'内容'"`
}

func (a AdminLog) TableName() string {
	return "admin_log"
}

func (a *AdminLog) FromPage(page, size int, order, address, ip, desc string) (data []AdminLog, total int) {
	if order == "" {
		order = "id desc"
	}
	where := "1=1"
	args := []any{}
	if address != "" {
		where += " AND `address` LIKE ?"
		args = append(args, "%"+address+"%")
	}
	if ip != "" {
		where += " AND `ip` LIKE ?"
		args = append(args, "%"+ip+"%")
	}
	if desc != "" {
		where += " AND `desc` LIKE ?"
		args = append(args, "%"+desc+"%")
	}
	subQuery := public.MainSql.Table(a.TableName()).Where(a).Where(where, args...)
	subQuery.Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	subQuery.Count(&total)
	return
}

func (a *AdminLog) Create() error {
	if a.Address == "" {
		a.Address = "未知地址"
	}
	a.Atime = time.Now()
	return public.MainSql.Table(a.TableName()).Create(a).Error
}
