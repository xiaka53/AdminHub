package dao

import (
	"github.com/xiaka53/AdminHub/public"
	"time"
)

type Admin struct {
	AdminId       uint      `gorm:"column:admin_id;primary_key;type:int(12) auto_increment;comment:'管理员Id'"`
	Username      string    `gorm:"column:username;type:varchar(10);not null;comment:'用户名'"`
	Password      string    `gorm:"column:password;type:varchar(60);not null;comment:'密码'"`
	LastLoginIp   string    `gorm:"column:last_login_ip;type:varchar(255);comment:'最后登陆IP'"`
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;comment:'最后登陆时间'"`
	RoleId        uint      `gorm:"column:role_id;type:int(12);not null;comment:'角色ID'"`
	Atime         time.Time `gorm:"column:atime;type:datetime;comment:'添加时间'"`
	Avatar        string    `gorm:"column:avatar;type:varchar(255);comment:'头像'"`
}

func (a Admin) TableName() string {
	return "admin"
}

func (a *Admin) First() error {
	return public.MainSql.Table(a.TableName()).Where(a).First(a).Error
}

func (a *Admin) Find() (data []Admin) {
	public.MainSql.Table(a.TableName()).Where(a).Find(&data)
	return
}

func (a *Admin) FindById() (data map[uint]Admin) {
	var _data []Admin
	public.MainSql.Table(a.TableName()).Where(a).Find(&_data)
	data = make(map[uint]Admin)
	for _, v := range _data {
		data[v.AdminId] = v
	}
	return
}

func (a *Admin) FromPage(page, size int, order string) (data []Admin, total int) {
	if order == "" {
		order = "admin_id asc"
	}
	public.MainSql.Table(a.TableName()).Where(a).Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	public.MainSql.Table(a.TableName()).Where(a).Count(&total)
	return
}

func (a *Admin) Edit() error {
	return public.MainSql.Table(a.TableName()).Where("admin_id=?", a.AdminId).Updates(a).Error
}

func (a *Admin) Create() error {
	a.Atime = time.Now()
	a.LastLoginTime = time.Now()
	return public.MainSql.Table(a.TableName()).Create(a).Error
}

func (a *Admin) Del() error {
	return public.MainSql.Table(a.TableName()).Delete(a).Error
}
