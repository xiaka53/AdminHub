package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/dao"
	"github.com/xiaka53/AdminHub/exec/dto"
	"github.com/xiaka53/AdminHub/middleware"
	"github.com/xiaka53/AdminHub/upload"
	"github.com/xiaka53/DeployAndLog/lib"
)

type setting struct {
}

func SettingRouterGroup(r *gin.RouterGroup) {
	var s setting
	r.POST("getUploadConfig", s.getUploadConfig)
	r.POST("saveQiniu", s.saveQiniu)
	r.POST("saveAlioss", s.saveAlioss)
	r.POST("saveTxcos", s.saveTxcos)
	r.POST("saveLocal", s.saveLocal)
	r.POST("settingList", s.settingList)
	r.POST("addSetting", s.addSetting)
	r.POST("editSetting", s.editSetting)
	r.POST("delSetting", s.delSetting)
	r.POST("getUploadToken", s.getUploadToken)
	r.POST("getFileList", s.getFileList)
	r.POST("addFile", s.addFile)
	r.POST("delFile", s.delFile)
}

func (setting) getUploadConfig(c *gin.Context) {
	var (
		info       uploadConfig_info
		_uploadSet dao.UploadSet
	)
	_ = (&_uploadSet).First()
	info = uploadConfig_info{
		Visible: int(_uploadSet.Visible),
	}
	_ = json.Unmarshal([]byte(_uploadSet.Qiniu), &info.Qiniu)
	_ = json.Unmarshal([]byte(_uploadSet.Alioss), &info.Alioss)
	_ = json.Unmarshal([]byte(_uploadSet.Txcos), &info.Txcos)
	middleware.ResponseSuccess(c, info)
}

type uploadConfig_info struct {
	Visible int `json:"visible"`
	Qiniu   any `json:"qiniu"`
	Alioss  any `json:"alioss"`
	Txcos   any `json:"txcos"`
}

func (setting) saveQiniu(c *gin.Context) {
	var (
		paramers   dto.SettingSaveQiniu
		_uploadSet dao.UploadSet
		qiniuByte  []byte
		err        error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	qiniuByte, err = json.Marshal(paramers.Qiniu)
	if err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if paramers.Visible == 1 {
		_uploadSet.Visible = dao.UploadSetVisibleQiniu
	}
	if _uploadSet.Visible == dao.UploadSetVisibleQiniu && paramers.Visible == 0 {
		_uploadSet.Visible = dao.UploadSetVisibleNotSelected
	}
	_uploadSet.Qiniu = string(qiniuByte)
	if err = (&_uploadSet).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) saveAlioss(c *gin.Context) {
	var (
		paramers    dto.SettingSaveAlioss
		_uploadSet  dao.UploadSet
		aliOrTxByte []byte
		err         error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	aliOrTxByte, err = json.Marshal(paramers.Ali)
	if err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if paramers.Visible == 2 {
		_uploadSet.Visible = dao.UploadSetVisibleAli
	}
	if _uploadSet.Visible == dao.UploadSetVisibleAli && paramers.Visible == 0 {
		_uploadSet.Visible = dao.UploadSetVisibleNotSelected
	}
	_uploadSet.Alioss = string(aliOrTxByte)
	if err = (&_uploadSet).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) saveTxcos(c *gin.Context) {
	var (
		paramers    dto.SettingSaveTxcos
		_uploadSet  dao.UploadSet
		aliOrTxByte []byte
		err         error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	aliOrTxByte, err = json.Marshal(paramers.Tx)
	if err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if paramers.Visible == 3 {
		_uploadSet.Visible = dao.UploadSetVisibleTx
	}
	if _uploadSet.Visible == dao.UploadSetVisibleTx && paramers.Visible == 0 {
		_uploadSet.Visible = dao.UploadSetVisibleNotSelected
	}
	_uploadSet.Txcos = string(aliOrTxByte)
	if err = (&_uploadSet).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) saveLocal(c *gin.Context) {
	var (
		paramers   dto.SettingSaveLocal
		_uploadSet dao.UploadSet
		err        error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if paramers.Visible == 4 {
		_uploadSet.Visible = dao.UploadSetVisibleLoc
	}
	if _uploadSet.Visible == dao.UploadSetVisibleLoc && paramers.Visible == 0 {
		_uploadSet.Visible = dao.UploadSetVisibleNotSelected
	}
	if err = (&_uploadSet).Edit(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) settingList(c *gin.Context) {
	var (
		paramers dto.SettingSettingList
		info     settingList_info
		data     []dao.Setting
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	data, info.All = (&dao.Setting{}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy)
	for _, v := range data {
		info.Datas = append(info.Datas, settingList_info_datas{
			Id:     v.Id,
			Title:  v.Title,
			Type:   v.Type,
			Value:  v.Value,
			CanDel: v.CanDel,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type settingList_info struct {
	All   int                      `json:"all"`
	Datas []settingList_info_datas `json:"datas"`
}

type settingList_info_datas struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Type   uint   `json:"type"`
	Value  string `json:"value"`
	CanDel uint   `json:"canDel"`
}

func (setting) addSetting(c *gin.Context) {
	var (
		paramers dto.SettingAddList
		_setting dao.Setting
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_setting.Title = paramers.Title
	if err = (&_setting).First(); err == nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	_setting = dao.Setting{
		Title:  paramers.Title,
		Type:   paramers.Type,
		Value:  fmt.Sprintf("%v", paramers.Value),
		CanDel: paramers.CanDel,
	}
	if err = (&_setting).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) editSetting(c *gin.Context) {
	var (
		paramers dto.SettingEditList
		_setting dao.Setting
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_setting.Id = paramers.Id
	if err = (&_setting).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	_setting.Title = paramers.Title
	_setting.Type = paramers.Type
	_setting.Value = fmt.Sprintf("%v", paramers.Value)
	if err = (&_setting).Edit(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) delSetting(c *gin.Context) {
	var (
		paramers dto.SettingDelList
		_setting dao.Setting
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_setting.Id = paramers.Id
	if err = (&_setting).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	if err = (&_setting).Del(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) getFileList(c *gin.Context) {
	var (
		paramers dto.SettingGetFileList
		info     fileList_info
		data     []dao.UploadFiles
		err      error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	info.Datas = []fileList_info_datas{}
	data, info.All = (&dao.UploadFiles{Pid: paramers.Pid, Type: dao.UploadFilesType(paramers.Type)}).FromPage(paramers.Page.Page, paramers.Size, paramers.OrderBy.OrderBy, paramers.Name)
	for _, v := range data {
		info.Datas = append(info.Datas, fileList_info_datas{
			Atime:  v.Atime.Format(lib.TimeFormat),
			Domain: uint(v.Domain),
			Id:     v.Id,
			Key:    v.Key,
			Name:   v.Name,
			Pid:    v.Pid,
			Type:   uint(v.Type),
			Url:    v.Url,
		})
	}
	middleware.ResponseSuccess(c, info)
}

type fileList_info struct {
	All   int                   `json:"all"`
	Datas []fileList_info_datas `json:"datas"`
}

type fileList_info_datas struct {
	Atime  string `json:"atime"`
	Domain uint   `json:"domain"`
	Id     uint   `json:"id"`
	Key    string `json:"key"`
	Name   string `json:"name"`
	Pid    uint   `json:"pid"`
	Type   uint   `json:"type"`
	Url    string `json:"url"`
}

func (setting) getUploadToken(c *gin.Context) {
	var (
		_uploadSet dao.UploadSet
		_upload    upload.UploadFile
		info       map[string]any
		err        error
	)
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	switch _uploadSet.Visible {
	case dao.UploadSetVisibleTx:
		_upload = upload.GetTx(_uploadSet.Txcos)
	case dao.UploadSetVisibleAli:
		_upload = upload.GetAli(_uploadSet.Alioss)
	case dao.UploadSetVisibleQiniu:
		_upload = upload.GetQiNiu(_uploadSet.Qiniu)
	case dao.UploadSetVisibleLoc:
		_upload = upload.GetLoc()
	default:
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	info = _upload.GetToken()
	info["visible"] = _uploadSet.Visible
	middleware.ResponseSuccess(c, info)
}

func (setting) addFile(c *gin.Context) {
	var (
		paramers     dto.SettingAddFile
		_uploadFiles dao.UploadFiles
		err          error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	_uploadFiles = dao.UploadFiles{
		Domain: dao.UploadFilesDomain(paramers.Domain),
		Type:   dao.UploadFilesType(paramers.Type),
		Name:   paramers.Name,
		Key:    paramers.Key,
		Url:    paramers.Url,
		Pid:    paramers.Pid,
	}
	if err = (&_uploadFiles).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func (setting) delFile(c *gin.Context) {
	var (
		paramers     dto.SettingDelFile
		_uploadSet   dao.UploadSet
		_upload      upload.UploadFile
		_uploadFiles dao.UploadFiles
		err          error
	)
	if err = (&paramers).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	if err = (&_uploadSet).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	_uploadFiles.Id = paramers.Id
	if err = (&_uploadFiles).First(); err != nil {
		middleware.ResponseError(c, middleware.EditErr, errors.New(""))
		return
	}
	if _uploadFiles.Type == dao.UploadFilesTypeFolder && (&dao.UploadFiles{Pid: _uploadFiles.Id}).Total() > 1 {
		middleware.ResponseError(c, middleware.FolderHaveFile, errors.New(""))
		return
	}
	switch _uploadFiles.Domain {
	case dao.UploadFilesDomainTx:
		_upload = upload.GetTx(_uploadSet.Txcos)
	case dao.UploadFilesDomainAli:
		_upload = upload.GetAli(_uploadSet.Alioss)
	case dao.UploadFilesDomainQiniu:
		_upload = upload.GetQiNiu(_uploadSet.Qiniu)
	case dao.UploadFilesDomainLoc:
		_upload = upload.GetLoc()
	default:
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	_upload.Delete(_uploadFiles.Key)
	if err = (&_uploadFiles).Del(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}
