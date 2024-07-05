package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
	"reflect"
)

type LoginLogin struct {
	Username string `form:"username" json:"username" validate:"min=5,max=16" zh:"用户名"`
	Password string `form:"password" json:"password" validate:"min=6,max=20" zh:"密码"`
	Code     string `form:"code" json:"code" validate:"len=4" zh:"验证码"`
	Uuid     string `form:"uuid" json:"uuid" validate:"omitempty,len=36" zh:"Uuid"`
}

func (o *LoginLogin) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
