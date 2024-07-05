package dao

import (
	"github.com/xiaka53/AdminHub/public"
	"time"
)

type Role struct {
	RoleId   uint      `gorm:"column:role_id;primary_key;type:int(12) auto_increment;comment:'角色Id'"`
	RoleName string    `gorm:"column:role_name;type:varchar(20);not null;comment:'角色名'"`
	Ids      string    `gorm:"column:ids;type:text;not null;comment:'权限组'"`
	Describe string    `gorm:"column:describe;type:varchar(255);comment:'描述'"`
	Atime    time.Time `gorm:"column:atime;type:datetime;comment:'添加时间'"`
}

func (r Role) TableName() string {
	return "role"
}

func (r *Role) First() error {
	return public.MainSql.Table(r.TableName()).Where(r).First(r).Error
}

func (r *Role) Find() (data []Role) {
	public.MainSql.Table(r.TableName()).Where(r).Find(&data)
	return
}

func (r *Role) FindById() map[uint]Role {
	var data []Role
	public.MainSql.Table(r.TableName()).Where(r).First(&data)
	_data := make(map[uint]Role)
	for _, v := range data {
		_data[v.RoleId] = v
	}
	return _data
}

func (r *Role) FromPage(page, size int, order string) (data []Role, total int) {
	if order == "" {
		order = "role_id asc"
	}
	public.MainSql.Table(r.TableName()).Where(r).Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	public.MainSql.Table(r.TableName()).Where(r).Count(&total)
	return
}

func (r *Role) Create() error {
	r.Atime = time.Now()
	return public.MainSql.Table(r.TableName()).Create(r).Error
}

func (r *Role) Edit() error {
	return public.MainSql.Table(r.TableName()).Where("role_id=?", r.RoleId).Updates(r).Error
}

func (r *Role) Del() error {
	return public.MainSql.Table(r.TableName()).Delete(r).Error
}
