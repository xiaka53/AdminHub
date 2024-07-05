package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/exec/dto"
	"github.com/xiaka53/AdminHub/middleware"
	"github.com/xiaka53/DeployAndLog/lib"
)

type role struct {
}

func RoleGroupRouter(r *gin.RouterGroup) {
	var _role role
	r.POST("addRoleGetMenus", _role.addRoleGetMenus)
	r.POST("roleList", _role.roleList)
	r.POST("editRole", _role.editRole)
	r.POST("addRole", _role.addRole)
	r.POST("delRole", _role.delRole)
}

func (role) roleList(c *gin.Context) {
	var (
		paramers dto.RoleRoleList
		info     roleList_info
		data     []dao.Role
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	data, info.All = (&dao.Role{}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy)
	for _, v := range data {
		info.Datas = append(info.Datas, roleList_info_data{
			Atime:    v.Atime.Format(lib.TimeFormat),
			Describe: v.Describe,
			Ids:      v.Ids,
			RoleId:   v.RoleId,
			RoleName: v.RoleName,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type roleList_info struct {
	All   int                  `json:"all"`
	Datas []roleList_info_data `json:"datas"`
}

type roleList_info_data struct {
	Atime    string `json:"atime"`
	Describe string `json:"describe"`
	Ids      string `json:"ids"`
	RoleId   uint   `json:"role_id"`
	RoleName string `json:"role_name"`
}

func (r role) addRoleGetMenus(c *gin.Context) {
	var (
		info      []addRoleGetMenus_info
		menuIndex = make(map[uint][]addRoleGetMenus_info)
	)
	for _, v := range (&dao.Menu{}).Find() {
		menuIndex[v.Pid] = append(menuIndex[v.Pid], addRoleGetMenus_info{
			Child: r.flashback(menuIndex[v.Id]),
			Id:    v.Id,
			Name:  v.Name,
		})
	}
	info = r.flashback(menuIndex[0])
	middleware.ResponseSuccess(c, info)
}

func (role) flashback(info []addRoleGetMenus_info) (data []addRoleGetMenus_info) {
	length := len(info)
	data = []addRoleGetMenus_info{}
	for i := 0; i < length; i++ {
		data = append(data, info[length-i-1])
	}
	return data
}

type addRoleGetMenus_info struct {
	Child []addRoleGetMenus_info `json:"child"`
	Id    uint                   `json:"id"`
	Name  string                 `json:"name"`
}

func (role) editRole(c *gin.Context) {
	var (
		paramers dto.RoleEditRole
		_role    dao.Role
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_role.RoleId = paramers.RoleId
	if err = (&_role).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	_role.Ids = paramers.Ids
	_role.RoleName = paramers.RoleName
	_role.Describe = paramers.Describe
	if err = (&_role).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (role) addRole(c *gin.Context) {
	var (
		paramers dto.RoleAddRole
		_role    dao.Role
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_role.Ids = paramers.Ids
	_role.RoleName = paramers.RoleName
	_role.Describe = paramers.Describe
	if err = (&_role).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (role) delRole(c *gin.Context) {
	var (
		paramers dto.RoleDelRole
		_admin   dao.Admin
		_role    dao.Role
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_role.RoleId = paramers.RoleId
	if err = (&_role).First(); err != nil {
		middleware.ResponseError(c, middleware.NoInfo, err)
		return
	}
	_admin.RoleId = _role.RoleId
	_ = (&_admin).First()
	if _admin.AdminId > 0 {
		middleware.ResponseError(c, middleware.RolehaveAdmin, err)
		return
	}
	if err = (&_role).Del(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
