package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/exec/dto"
	"github.com/xiaka53/AdminHub/middleware"
	"github.com/xiaka53/DeployAndLog/lib"
	"golang.org/x/crypto/bcrypt"
)

type admin struct {
}

func AdminRouterGroup(r *gin.RouterGroup) {
	var a admin
	r.POST("getLoginInfo", a.getLoginInfo)
	r.POST("getSearchRoleList", a.getSearchRoleList)
	r.POST("adminList", a.adminList)
	r.POST("addAdmin", a.addAdmin)
	r.POST("editAdmin", a.editAdmin)
	r.POST("delAdmin", a.delAdmin)
	r.POST("editPwd", a.editPwd)
	r.POST("editAvatar", a.editAvatar)
	r.POST("getSearchAdminList", a.getSearchAdminList)
	r.POST("adminLog", a.adminLog)
}

func (a admin) getLoginInfo(c *gin.Context) {
	var (
		info       loginInfo_info
		_admin     dao.Admin
		_setting   dao.Setting
		menuIndexK int
		menuIndex  = make(map[uint]int)
		menuChild  = make(map[uint][]loginInfo_info_menus_child)
		ids        any
		ok         bool
	)
	_setting.Title = dao.SystemName
	_ = (&_setting).First()
	info.Name = _setting.Value
	_admin.AdminId = c.GetUint("_admin_id_")
	if err := (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.LoginErr, errors.New(""))
		return
	}
	ids, ok = c.Get("_roleIds_")
	if !ok {
		ids = []uint{}
	}
	info.Username = _admin.Username
	info.Avatar = _admin.Avatar
	info.Menus = []loginInfo_info_menus{}
	for _, v := range (&dao.Menu{}).GetMenu(ids.([]uint)) {
		if v.Level == 1 {
			info.Menus = append(info.Menus, loginInfo_info_menus{
				Child:                      nil,
				loginInfo_info_menus_child: loginInfo_info_menus_child{v.Icon, v.Path, v.Name, v.Id},
			})
			menuIndex[v.Id] = menuIndexK
			menuIndexK++
		}
		if v.Level == 2 {
			menuChild[v.Pid] = append(menuChild[v.Pid], loginInfo_info_menus_child{v.Icon, v.Path, v.Name, v.Id})
		}
	}
	for pid, child := range menuChild {
		if k, ok := menuIndex[pid]; ok {
			info.Menus[k].Child = child
		}
	}
	middleware.ResponseSuccess(c, info)
}

type loginInfo_info struct {
	Avatar   string                 `json:"avatar"`
	Name     string                 `json:"name"`
	Username string                 `json:"username"`
	Menus    []loginInfo_info_menus `json:"menus"`
}

type loginInfo_info_menus struct {
	Child []loginInfo_info_menus_child `json:"child"`
	loginInfo_info_menus_child
}

type loginInfo_info_menus_child struct {
	Icon  string `json:"icon"`
	Path  string `json:"path"`
	Title string `json:"title"`
	Id    uint   `json:"id"`
}

func (admin) getSearchRoleList(c *gin.Context) {
	var (
		info []searchRoleList_info
	)
	for _, v := range (&dao.Role{}).Find() {
		info = append(info, searchRoleList_info{
			Id:   v.RoleId,
			Name: v.RoleName,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type searchRoleList_info struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (admin) adminList(c *gin.Context) {
	var (
		paramers dto.AdminAdminList
		info     adminList_info
		roles    map[uint]dao.Role
		roleId   uint
		data     []dao.Admin
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	if paramers.RoleId != "" {
		roleId = uint(paramers.RoleId.(float64))
	}
	roles = (&dao.Role{}).FindById()
	data, info.All = (&dao.Admin{Username: paramers.Username, RoleId: roleId}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy)
	info.Datas = []adminList_info_datas{}
	for _, v := range data {
		sun := adminList_info_datas{
			AdminId:       v.AdminId,
			Atime:         v.Atime.Format(lib.TimeFormat),
			LastLoginIp:   v.LastLoginIp,
			LastLoginTime: v.LastLoginTime.Format(lib.TimeFormat),
			RoleId:        v.RoleId,
			RoleName:      roles[v.RoleId].RoleName,
			Username:      v.Username,
		}
		if sun.Atime == sun.LastLoginTime {
			sun.LastLoginTime = ""
		}
		info.Datas = append(info.Datas, sun)
	}
	middleware.ResponseSuccess(c, info)
}

type adminList_info struct {
	All   int
	Datas []adminList_info_datas `json:"datas"`
}

type adminList_info_datas struct {
	AdminId       uint   `json:"admin_id"`
	Atime         string `json:"atime"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime string `json:"last_login_time"`
	RoleId        uint   `json:"role_id"`
	RoleName      string `json:"role_name"`
	Username      string `json:"username"`
}

func (admin) addAdmin(c *gin.Context) {
	var (
		paramers dto.AdminAddAdmin
		_role    dao.Role
		_admin   dao.Admin
		password []byte
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_role.RoleId = paramers.RoleId
	if err = (&_role).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	_admin.RoleId = _role.RoleId
	_admin.Username = paramers.Username
	if paramers.Password == "" {
		paramers.Password = "123456"
	}
	password, err = bcrypt.GenerateFromPassword([]byte(paramers.Password), bcrypt.DefaultCost)
	_admin.Password = string(password)
	if err = (&_admin).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (admin) editAdmin(c *gin.Context) {
	var (
		paramers dto.AdminEditAdmin
		_role    dao.Role
		_admin   dao.Admin
		password []byte
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_role.RoleId = paramers.RoleId
	if err = (&_role).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	_admin.AdminId = paramers.AdminId
	if err = (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	_admin.RoleId = _role.RoleId
	_admin.Username = paramers.Username
	if paramers.Password == "" {
		paramers.Password = "123456"
	}
	password, err = bcrypt.GenerateFromPassword([]byte(paramers.Password), bcrypt.DefaultCost)
	_admin.Password = string(password)
	if err = (&_admin).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (admin) delAdmin(c *gin.Context) {
	var (
		paramers dto.AdminDelAdmin
		_admin   dao.Admin
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}

	_admin.AdminId = paramers.AdminId
	if err = (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if err = (&_admin).Del(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (admin) editPwd(c *gin.Context) {
	var (
		paramers dto.AdminEditPwd
		_admin   dao.Admin
		password []byte
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_admin.AdminId = c.GetUint("_admin_id_")
	if _admin.AdminId < 1 {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if err = (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(_admin.Password), []byte(paramers.OldPwd))
	if err != nil {
		middleware.ResponseError(c, middleware.PassErr, errors.New(""))
		return
	}
	password, err = bcrypt.GenerateFromPassword([]byte(paramers.Password), bcrypt.DefaultCost)
	_admin.Password = string(password)
	if err = (&_admin).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (admin) editAvatar(c *gin.Context) {
	var (
		paramers dto.AdminEditAvatar
		_admin   dao.Admin
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_admin.AdminId = c.GetUint("_admin_id_")
	if _admin.AdminId < 1 {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if err = (&_admin).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	_admin.Avatar = paramers.Avatar
	_admin.Username = paramers.Username
	if err = (&_admin).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (admin) getSearchAdminList(c *gin.Context) {
	var (
		info []searchAdminList_info
	)
	for _, v := range (&dao.Admin{}).Find() {
		info = append(info, searchAdminList_info{
			Name: v.Username,
			Id:   v.AdminId,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type searchAdminList_info struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

func (admin) adminLog(c *gin.Context) {
	var (
		paramers  dto.AdminAdminLogList
		info      adminLog_info
		data      []dao.AdminLog
		adminData map[uint]dao.Admin
		adminId   uint
		err       error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	if paramers.AdminId != "" {
		adminId = uint(paramers.AdminId.(float64))
	}
	data, info.All = (&dao.AdminLog{AdminId: adminId}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy, paramers.Address, paramers.Ip, paramers.Desc)
	info.Datas = []adminLog_info_datas{}
	if len(data) > 0 {
		adminData = (&dao.Admin{}).FindById()
	}
	for _, v := range data {
		info.Datas = append(info.Datas, adminLog_info_datas{
			Address:  v.Address,
			Atime:    v.Atime.Format(lib.TimeFormat),
			Desc:     v.Desc,
			Ip:       v.Ip,
			Username: adminData[v.AdminId].Username,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type adminLog_info struct {
	All   int                   `json:"all"`
	Datas []adminLog_info_datas `json:"datas"`
}

type adminLog_info_datas struct {
	Address  string `json:"address"`
	Atime    string `json:"atime"`
	Desc     string `json:"desc"`
	Ip       string `json:"ip"`
	Username string `json:"username"`
}
