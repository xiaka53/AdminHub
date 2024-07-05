package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/exec/dto"
	"github.com/xiaka53/AdminHub/middleware"
)

type menu struct {
}

func MenuRouterGroup(r *gin.RouterGroup) {
	var m menu
	r.POST("menuList", m.menuList)
	r.POST("addMenu", m.addMenu)
	r.POST("editMenu", m.editMenu)
	r.POST("delMenu", m.delMenu)
	r.POST("setNeedLog", m.setNeedLog)
	r.POST("getMenusByPid", m.getMenusByPid)
}

func (m menu) menuList(c *gin.Context) {
	var (
		paramers dto.MenuMenuList
		info     menuList_info
		data     []dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	data, info.All = (&dao.Menu{Level: 1}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy)
	for _, v := range data {
		info.Datas = append(info.Datas, menuList_info_datas{m.getMenuListChild(v.Id), v})
	}
	middleware.ResponseSuccess(c, info)
}

func (m menu) getMenuListChild(mId uint) []menuList_info_datas {
	child := (&dao.Menu{Pid: mId}).FindChild()
	data := []menuList_info_datas{}
	for _, v := range child {
		data = append(data, menuList_info_datas{
			Child: m.getMenuListChild(v.Id),
			Menu:  v,
		})
	}
	return data
}

type menuList_info struct {
	Datas []menuList_info_datas `json:"datas"`
	All   int                   `json:"all"`
}

type menuList_info_datas struct {
	Child []menuList_info_datas `json:"child"`
	dao.Menu
}

func (menu) addMenu(c *gin.Context) {
	var (
		paramers dto.MenuAddMenu
		_menu    dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_menu = dao.Menu{
		Pid:     paramers.Pid,
		Name:    paramers.Name,
		Path:    paramers.Path,
		Route:   paramers.Route,
		Level:   paramers.Level,
		Icon:    paramers.Icon,
		Sort:    paramers.Sort,
		NeedLog: dao.MenuNeedLog(paramers.NeedLog),
		Display: dao.MenuDisplay(paramers.Display),
	}
	if err = (&_menu).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (menu) editMenu(c *gin.Context) {
	var (
		paramers dto.MenuEditMenu
		_menu    dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_menu.Id = paramers.Id
	if err = (&_menu).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	_menu = dao.Menu{
		Id:      paramers.Id,
		Pid:     paramers.Pid,
		Name:    paramers.Name,
		Path:    paramers.Path,
		Route:   paramers.Route,
		Level:   paramers.Level,
		Icon:    paramers.Icon,
		Sort:    paramers.Sort,
		NeedLog: dao.MenuNeedLog(paramers.NeedLog),
		Display: dao.MenuDisplay(paramers.Display),
	}
	if err = (&_menu).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (menu) delMenu(c *gin.Context) {
	var (
		paramers dto.MenuDelMenu
		_menu    dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_menu.Id = paramers.Id
	if err = (&_menu).First(); err != nil {
		middleware.ResponseError(c, middleware.NoInfo, err)
		return
	}
	if err = (&_menu).Del(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (menu) setNeedLog(c *gin.Context) {
	var (
		paramers dto.MenuSetNeedLog
		_menu    dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_menu.Id = paramers.Id
	if err = (&_menu).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	_menu.NeedLog = dao.MenuNeedLog(paramers.NeedLog)
	if err = (&_menu).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (menu) getMenusByPid(c *gin.Context) {
	var (
		paramers dto.MenuGetMenusByPid
		data     []dao.Menu
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	data = (&dao.Menu{Pid: paramers.Pid}).Find()
	middleware.ResponseSuccess(c, data)
}
