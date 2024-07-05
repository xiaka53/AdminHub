package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
	"reflect"
)

type RoleRoleList struct {
	OrderBy
	Page
}

func (o *RoleRoleList) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type RoleAddRole struct {
	Describe string `form:"describe" json:"describe" validate:"min=1" zh:"角色描述"`
	Ids      string `form:"ids" json:"ids" validate:"min=1" zh:"角色权限"`
	RoleName string `form:"Role_name" json:"Role_name" validate:"min=1,max=20" zh:"角色名称"`
}

func (o *RoleAddRole) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type RoleDelRole struct {
	RoleId uint `form:"role_id" json:"role_id" validate:"min=1" zh:"角色ID"`
}

func (o *RoleDelRole) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type RoleEditRole struct {
	RoleAddRole
	RoleDelRole
}

func (o *RoleEditRole) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
