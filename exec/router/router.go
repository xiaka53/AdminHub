package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/exec/api"
	"github.com/xiaka53/AdminHub/middleware"
	"net/http"
)

type Roter struct {
	Group map[string]map[string]func(group *gin.RouterGroup)
	R     *gin.Engine
}

func CreateRouter() *Roter {
	r := &Roter{
		Group: make(map[string]map[string]func(group *gin.RouterGroup)),
	}
	r.InitRouter()
	return r
}

func (r *Roter) InitRouter() {
	r.R = gin.New()
	r.R.Use(gin.Recovery())

	r.R.Use(middleware.RecoverMiddleware(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	r.R.Use(middleware.AccessMiddleware()) // 跨域
	r.R.StaticFS("/files", http.Dir("./files"))
	r.R.POST("upload", api.Upload)
	r.R.Use(middleware.AuthMiddleware())
	adminGroup := r.R.Group("admin")
	api.LoginRouterGroup(adminGroup.Group("login"))
	adminGroup.Use(middleware.LoginMiddleware())
	adminGroup.Use(middleware.LogMiddleware())
	api.AdminRouterGroup(adminGroup.Group("admin"))
	adminGroup.Use(middleware.RequestLog()) // 请求日志中间件
	api.RoleGroupRouter(adminGroup.Group("role"))
	api.MenuRouterGroup(adminGroup.Group("menu"))
	api.SettingRouterGroup(adminGroup.Group("setting"))
	return
}

func (r *Roter) SetGroup(groupName, funcGroupName string, _func func(group *gin.RouterGroup)) {
	if group, ok := r.Group[groupName]; ok {
		group[funcGroupName] = _func
		r.Group[groupName] = group
	} else {
		group = make(map[string]func(group *gin.RouterGroup))
		group[funcGroupName] = _func
		r.Group[groupName] = group
	}
}

func (r *Roter) Write() {
	for name, groupFunc := range r.Group {
		_group := r.R.Group(name)
		for funcGroupName, _func := range groupFunc {
			_func(_group.Group(funcGroupName))
		}
	}
}

// 路由初始化
func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Recovery())

	r.Use(middleware.RecoverMiddleware(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	r.Use(middleware.AccessMiddleware()) // 跨域
	r.StaticFS("/files", http.Dir("./files"))
	r.POST("upload", api.Upload)
	r.Use(middleware.AuthMiddleware())
	adminGroup := r.Group("admin")
	api.LoginRouterGroup(adminGroup.Group("login"))
	adminGroup.Use(middleware.LoginMiddleware())
	adminGroup.Use(middleware.LogMiddleware())
	api.AdminRouterGroup(adminGroup.Group("admin"))
	adminGroup.Use(middleware.RequestLog()) // 请求日志中间件
	api.RoleGroupRouter(adminGroup.Group("role"))
	api.MenuRouterGroup(adminGroup.Group("menu"))
	api.SettingRouterGroup(adminGroup.Group("setting"))
	return
}
