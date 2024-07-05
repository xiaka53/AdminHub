package dao

import (
	"github.com/xiaka53/AdminHub/public"
)

type MenuDisplay uint
type MenuNeedLog uint

const (
	MenuNeedLogN MenuNeedLog = 0 // 不写日志
	MenuNeedLogY MenuNeedLog = 1 // 写日志
	MenuDisplayY MenuDisplay = 1 // 展示
	MenuDisplayN MenuDisplay = 0 // 不展示
)

type Menu struct {
	Id      uint        `json:"id" gorm:"column:id;primaryKey;type:int(12) auto_increment;comment:'ID'"`
	Pid     uint        `json:"pid" gorm:"column:pid;type:int(12);comment:'上级ID'"`
	Name    string      `json:"name" gorm:"column:name;type:varchar(64);comment:'菜单名称'"`
	Path    string      `json:"path" gorm:"column:path;type:varchar(64);comment:'前端路由'"`
	Route   string      `json:"route" gorm:"column:route;type:varchar(64);comment:'后端路由'"`
	Level   uint        `json:"level" gorm:"column:level;type:tinyint(1);comment:'菜单等级'"`
	Icon    string      `json:"icon" gorm:"column:icon;type:varchar(32);not null;comment:'图标'"`
	Sort    uint        `json:"sort" gorm:"column:sort;type:int(12);comment:'排序'"`
	NeedLog MenuNeedLog `json:"needLog" gorm:"column:needLog;type:tinyint(1);comment:'是否需要日志'"`
	Display MenuDisplay `json:"display" gorm:"column:display;type:tinyint(1);comment:'是否显示'"`
}

func (m Menu) TableName() string {
	return "menu"
}

func (m *Menu) GetMenu(ids []uint) (data []Menu) {
	public.MainSql.Table(m.TableName()).Where("level<3 and display=? and id in (?)", MenuDisplayY, ids).Order("sort asc").Find(&data)
	return
}

func (m *Menu) First() error {
	return public.MainSql.Table(m.TableName()).Where(m).First(m).Error
}

func (m *Menu) Find() (data []Menu) {
	public.MainSql.Table(m.TableName()).Where(m).Order("`level` desc").Order("sort desc").Find(&data)
	return
}

func (m *Menu) FindChild() (data []Menu) {
	public.MainSql.Table(m.TableName()).Where(m).Order("sort asc").Find(&data)
	return
}

func (m *Menu) FromPage(page, size int, order string) (data []Menu, total int) {
	if order == "" {
		order = "sort asc"
	}
	public.MainSql.Table(m.TableName()).Where(m).Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	public.MainSql.Table(m.TableName()).Where(m).Count(&total)
	return
}

func (m *Menu) Create() error {
	return public.MainSql.Table(m.TableName()).Create(m).Error
}

func (m *Menu) Edit() error {
	return public.MainSql.Table(m.TableName()).Where("id=?", m.Id).Save(m).Error
}

func (m *Menu) Del() error {
	return public.MainSql.Table(m.TableName()).Delete(m).Error
}
