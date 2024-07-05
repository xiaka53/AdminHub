package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
	"reflect"
)

type AdminAdminList struct {
	OrderBy
	Page
	Username string `form:"username" json:"username" validate:"omitempty,min=1" zh:"用户名"`
	RoleId   any    `form:"role_id" json:"role_id" validate:"omitempty" zh:"角色ID"`
}

func (o *AdminAdminList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminAddAdmin struct {
	Password string `form:"password" json:"password" validate:"omitempty,min=6,max=20" zh:"密码"`
	Username string `form:"username" json:"username" validate:"min=5,max=16" zh:"用户名"`
	RoleId   uint   `form:"role_id" json:"role_id" validate:"min=1" zh:"角色ID"`
}

func (o *AdminAddAdmin) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminEditAdmin struct {
	AdminDelAdmin
	AdminAddAdmin
}

func (o *AdminEditAdmin) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminDelAdmin struct {
	AdminId uint `form:"admin_id" json:"admin_id" validate:"min=1" zh:"用户Id"`
}

func (o *AdminDelAdmin) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminEditPwd struct {
	OldPwd   string `form:"oldPwd" json:"oldPwd" validate:"min=6,max=20" zh:"原密码"`
	Password string `form:"password" json:"password" validate:"min=6,max=20" zh:"新密码"`
	Pwd1     string `form:"pwd1" json:"pwd1" validate:"eqfield=Password" zh:"确认密码密码"`
}

func (o *AdminEditPwd) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminEditAvatar struct {
	Avatar   string `form:"oldPwd" json:"avatar" validate:"omitempty,url" zh:"头像"`
	Username string `form:"username" json:"username" validate:"min=5,max=16" zh:"昵称"`
}

func (o *AdminEditAvatar) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type AdminAdminLogList struct {
	OrderBy
	Page
	Address string `form:"address" json:"address" validate:"omitempty,min=1" zh:"操作内容"`
	Desc    string `form:"desc" json:"desc" validate:"omitempty,min=1" zh:"操作内容"`
	Ip      string `form:"ip" json:"ip" validate:"omitempty,min=1" zh:"ip"`
	AdminId any    `form:"admin_id" json:"admin_id" validate:"omitempty" zh:"管理员ID"`
}

func (o *AdminAdminLogList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
