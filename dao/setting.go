package dao

import (
	"github.com/xiaka53/AdminHub/public"
)

type SettingType uint8
type SettingCanDel uint8

const (
	SettingTypeText  SettingType = 1 // 文本
	SettingTypeImage SettingType = 2 // 图片

	SettingCanDelUnAllow SettingType = 0 // 不允许
	SettingCanDelAllow   SettingType = 1 // 允许

	SystemName = "系统名称" // 系统名称
)

type Setting struct {
	Id     uint   `gorm:"column:id;primaryKey;type:int(12) auto_increment;comment:'ID'"`
	Title  string `gorm:"column:title;type:varchar(64);not null;comment:'配置标题'"`
	Type   uint   `gorm:"column:type;type:tinyint;not null;comment:'类型'"`
	Value  string `gorm:"column:value;type:text;not null;comment:'属性值'"`
	CanDel uint   `gorm:"column:canDel;type:tinyint;not null;comment:'是否允许删除'"`
}

func (s Setting) TableName() string {
	return "setting"
}

func (s *Setting) First() error {
	return public.MainSql.Table(s.TableName()).Where(s).First(s).Error
}

func (s *Setting) FromPage(page, size int, order string) (data []Setting, total int) {
	if order == "" {
		order = "id asc"
	}
	public.MainSql.Table(s.TableName()).Where(s).Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	public.MainSql.Table(s.TableName()).Where(s).Count(&total)
	return
}

func (s *Setting) Create() error {
	return public.MainSql.Table(s.TableName()).Create(s).Error
}

func (s *Setting) Edit() error {
	return public.MainSql.Table(s.TableName()).Where("id=?", s.Id).Updates(s).Error
}

func (s *Setting) Del() error {
	return public.MainSql.Table(s.TableName()).Delete(s).Error
}
