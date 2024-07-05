package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/public"
	"reflect"
)

type SettingSaveQiniu struct {
	dao.Qiniu
	SettingSaveLocal
}

func (o *SettingSaveQiniu) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingSaveAlioss struct {
	dao.Ali
	SettingSaveLocal
}

func (o *SettingSaveAlioss) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingSaveTxcos struct {
	dao.Tx
	SettingSaveLocal
}

func (o *SettingSaveTxcos) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingSaveLocal struct {
	Visible uint `form:"visible" json:"visible" validate:"omitempty,oneof=0 1 2 3 4" zh:"展示状态"`
}

func (o *SettingSaveLocal) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingSettingList struct {
	OrderBy
	Page
}

func (o *SettingSettingList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingAddList struct {
	CanDel uint   `form:"canDel" json:"canDel" validate:"oneof=0 1" zh:"允许删除"`
	Title  string `form:"title" json:"title" validate:"min=1" zh:"配置名称"`
	Type   uint   `form:"type" json:"type" validate:"oneof=0 1 2 3 4 5" zh:"类型"`
	Value  any    `form:"value" json:"value" validate:"omitempty" zh:"值"`
}

func (o *SettingAddList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingEditList struct {
	Title string `form:"title" json:"title" validate:"min=1" zh:"配置名称"`
	Type  uint   `form:"type" json:"type" validate:"oneof=0 1 2 3 4 5" zh:"类型"`
	Value any    `form:"value" json:"value" validate:"omitempty" zh:"值"`
	SettingDelList
}

func (o *SettingEditList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingDelList struct {
	Id uint `form:"id" json:"id" validate:"min=1" zh:"ID"`
}

func (o *SettingDelList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingGetFileList struct {
	OrderBy
	Page
	Type uint   `form:"type" json:"type" validate:"oneof=0 1 2 3 4 5 6 7" zh:"类型"`
	Pid  uint   `form:"pid" json:"pid" validate:"min=0" zh:"层级"`
	Name string `form:"name" json:"name" validate:"omitempty" zh:"名称"`
}

func (o *SettingGetFileList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingAddFile struct {
	Domain uint   `form:"domain" json:"domain" validate:"oneof=0 1 2 3" zh:"保存方式"`
	Key    string `form:"key" json:"key" validate:"min=0" zh:"key"`
	Name   string `form:"name" json:"name" validate:"min=1" zh:"名称"`
	Pid    uint   `form:"pid" json:"pid" validate:"min=0" zh:"位置"`
	Type   uint   `form:"type" json:"type" validate:"oneof=1 2 3 4 5 6 7 8" zh:"类型"`
	Url    string `form:"url" json:"url" validate:"min=0" zh:"访问链接"`
}

func (o *SettingAddFile) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type SettingDelFile struct {
	Id uint `form:"id" json:"id" validate:"min=1" zh:"ID"`
}

func (o *SettingDelFile) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
