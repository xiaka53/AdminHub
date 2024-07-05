package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
	"reflect"
)

type MenuMenuList struct {
	OrderBy
	Page
}

func (o *MenuMenuList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type MenuAddMenu struct {
	Display int    `form:"display" json:"display" validate:"min=0" zh:"是否显示"`
	Icon    string `form:"icon" json:"icon" validate:"omitempty,min=1" zh:"图标"`
	Level   uint   `form:"level" json:"level" validate:"min=1" zh:"菜单等级"`
	Name    string `form:"name" json:"name" validate:"min=1,max=20" zh:"菜单名称"`
	NeedLog int    `form:"needLog" json:"needLog" validate:"min=0" zh:"是否展示日志"`
	Path    string `form:"path" json:"path" validate:"omitempty,min=1" zh:"前端路由"`
	Pid     uint   `form:"pid" json:"pid" validate:"min=0" zh:"上级"`
	Route   string `form:"route" json:"route" validate:"omitempty,min=1" zh:"后端路由"`
	Sort    uint   `form:"sort" json:"sort" validate:"min=0" zh:"排序"`
}

func (o *MenuAddMenu) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type MenuDelMenu struct {
	Id uint `form:"id" json:"id" validate:"min=1" zh:"菜单ID"`
}

func (o *MenuDelMenu) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type MenuEditMenu struct {
	MenuAddMenu
	MenuDelMenu
}

func (o *MenuEditMenu) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type MenuSetNeedLog struct {
	NeedLog int `form:"needLog" json:"needLog" validate:"min=0" zh:"是否展示日志"`
	MenuDelMenu
}

func (o *MenuSetNeedLog) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type MenuGetMenusByPid struct {
	Pid uint `form:"pid" json:"pid" validate:"min=0" zh:"上级ID"`
}

func (o *MenuGetMenusByPid) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
