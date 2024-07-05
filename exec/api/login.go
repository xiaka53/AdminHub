package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/exec/dto"
	"github.com/xiaka53/AdminHub/middleware"
	"github.com/xiaka53/AdminHub/public"
	"github.com/xiaka53/AdminHub/redis"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type login struct {
}

func LoginRouterGroup(r *gin.RouterGroup) {
	var l login
	r.POST("login", l.login, middleware.LogMiddleware())
	r.POST("getCaptcha", l.getCaptcha)
	r.POST("getSystemName", l.getSystemName)
}

func (login) getSystemName(c *gin.Context) {
	var (
		info     systemNameInfo
		_setting dao.Setting
	)
	_setting.Title = dao.SystemName
	_ = (&_setting).First()
	info.Name = _setting.Value
	middleware.ResponseSuccess(c, info)
}

type systemNameInfo struct {
	Name string `json:"name"`
}

func (login) getCaptcha(c *gin.Context) {
	var (
		info captchaInfo
		_l   redis.Login
	)
	info.Img, _l.Code = public.GetCodeImage()
	info.Uuid = public.GetUUid()
	_l.Uuid = info.Uuid
	(&_l).SetUuid()
	middleware.ResponseSuccess(c, info)
}

type captchaInfo struct {
	Img  string `json:"img"`
	Uuid string `json:"uuid"`
}

func (login) login(c *gin.Context) {
	var (
		paramers dto.LoginLogin
		info     login_info
		_l       redis.Login
		_admin   dao.Admin
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	if paramers.Uuid == "" {
		middleware.ResponseError(c, middleware.CodeErr, errors.New(""))
		return
	}
	_l.Uuid = paramers.Uuid
	(&_l).GetCode()
	if strings.ToLower(paramers.Code) != strings.ToLower(_l.Code) {
		middleware.ResponseError(c, middleware.CodeErr, errors.New(""))
		return
	}
	_admin.Username = paramers.Username
	if err = (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.UserOrPassErr, errors.New(""))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(_admin.Password), []byte(paramers.Password))
	if err != nil {
		middleware.ResponseError(c, middleware.UserOrPassErr, errors.New(""))
		return
	}
	_l.AdminId = _admin.AdminId
	(_l).SetToken()
	info = login_info{
		Avatar:   _admin.Avatar,
		NickName: _admin.Username,
		Token:    _l.Token,
	}
	c.Set("_admin_id_", _l.AdminId)
	c.Set("_auth_menu_log_", true)
	c.Set("_auth_menu_log_name_", "登陆")
	c.Set("_auth_menu_log_auth_", "Login")
	middleware.ResponseSuccess(c, info)
}

type login_info struct {
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
	Token    string `json:"token"`
}
