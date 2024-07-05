package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/public"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()
		menu := dao.Menu{Route: fullPath}
		_ = (&menu).First()
		c.Set("_auth_menu_id_", menu.Id)
		if menu.NeedLog == dao.MenuNeedLogY {
			c.Set("_auth_menu_log_", true)
			c.Set("_auth_menu_log_name_", menu.Name)
			c.Set("_auth_menu_log_auth_", fullPath)
		}
		c.Next()
	}
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLog := c.GetBool("_auth_menu_log_")
		if isLog {
			_ = (&dao.AdminLog{
				AdminId: c.GetUint("_admin_id_"),
				Ip:      c.ClientIP(),
				Address: public.GetIPInfo(c.ClientIP()).Address,
				Route:   c.GetString("_auth_menu_log_auth_"),
				Desc:    c.GetString("_auth_menu_log_name_"),
			}).Create()
		}
		c.Next()
	}
}
