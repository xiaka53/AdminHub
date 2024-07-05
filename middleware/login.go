package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/redis"
	"strconv"
	"strings"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _l redis.Login
		_l.Token = c.GetHeader("token")
		if _l.Token == "" {
			ResponseError(c, LoginErr, errors.New(""))
			return
		}
		(&_l).GetToken()
		if _l.AdminId < 1 {
			ResponseError(c, LoginErr, errors.New(""))
			return
		}
		admin := dao.Admin{AdminId: _l.AdminId}
		if err := (&admin).First(); err != nil {
			ResponseError(c, LoginErr, errors.New(""))
			return
		}
		role := dao.Role{RoleId: admin.RoleId}
		if err := (&role).First(); err != nil {
			ResponseError(c, RoleErr, errors.New(""))
			return
		}
		if len(role.Ids) < 2 {
			ResponseError(c, RoleErr, errors.New(""))
			return
		}
		var roleIds []uint
		for _, v := range strings.Split(role.Ids[1:len(role.Ids)-1], ",") {
			rId, _ := strconv.Atoi(v)
			if rId == 0 {
				continue
			}
			roleIds = append(roleIds, uint(rId))
		}
		c.Set("_roleIds_", roleIds)
		menuId := c.GetUint("_auth_menu_id_")
		if menuId > 0 {
			var ok bool
			for _, v := range roleIds {
				if v == menuId {
					ok = true
					break
				}
			}
			if !ok {
				ResponseError(c, RoleErr, errors.New(""))
				return
			}
		}
		c.Set("_admin_id_", _l.AdminId)
		c.Next()
	}
}
