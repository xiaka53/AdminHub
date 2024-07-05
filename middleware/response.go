package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/DeployAndLog/lib"
)

type ResponseCode int

var (
	SuccessCode       ResponseCode = 1
	InternalErrorCode ResponseCode = 500 // 系统内部错误

	LoginErr       ResponseCode = 999  // 登陆错误!
	RoleErr        ResponseCode = 1000 // 权限错误!
	CodeErr        ResponseCode = 1001 // 验证码错误
	UserOrPassErr  ResponseCode = 1002 // 用户名或者密码错误!
	EditErr        ResponseCode = 1003 // 修改失败!
	AddErr         ResponseCode = 1004 // 添加失败!
	NoInfo         ResponseCode = 1005 // 不存在的信息!
	RolehaveAdmin  ResponseCode = 1006 // 该角色组有管理员!
	PassErr        ResponseCode = 1007 // 密码错误!
	FolderHaveFile ResponseCode = 1008 // 文件夹內有文件!
	ParameterError ResponseCode = 3001 // 参数错误
)

var zh, en *ResponseMsgLang

func init() {
	zh = createResponseMsgLang("zh")
	zh.setMessage(map[ResponseCode]string{
		SuccessCode:       "操作成功",
		InternalErrorCode: "系统内部错误",
		LoginErr:          "登陆错误！",
		RoleErr:           "权限错误！",
		CodeErr:           "验证码错误",
		UserOrPassErr:     "用户名或者密码错误！",
		EditErr:           "修改失败！",
		AddErr:            "添加失败！",
		NoInfo:            "不存在的信息！",
		RolehaveAdmin:     "该角色组有管理员！",
		PassErr:           "密码错误！",
		FolderHaveFile:    "文件夹內有文件！",
		ParameterError:    "参数错误！",
	})
	zh.setDefacode(zh.getMessage(InternalErrorCode))
	en = createResponseMsgLang("en")
	en.setMessage(map[ResponseCode]string{
		SuccessCode:       "success",
		InternalErrorCode: "系统内部错误",
		LoginErr:          "Internal system error！",
		RoleErr:           "Permission error！",
		CodeErr:           "Verification code error",
		UserOrPassErr:     "Incorrect username or password！",
		EditErr:           "fail to edit！",
		AddErr:            "add failed！",
		NoInfo:            "No information！",
		RolehaveAdmin:     "The role group has an administrator！",
		PassErr:           "wrong password！",
		FolderHaveFile:    "There are files in the folder！",
		ParameterError:    "Parameter error！",
	})
	en.setDefacode(en.getMessage(InternalErrorCode))
}

// 返回信息格式
type Response struct {
	Code    ResponseCode `json:"code"`
	Msg     string       `json:"msg"`
	Data    interface{}  `json:"data"`
	TraceId interface{}  `json:"trace_id"`
}

// 返回前端错误信息
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	var (
		trace        interface{}
		traceContext *lib.TraceContext
		traceId      string
		resp         Response
		respone      []byte
	)
	trace, _ = c.Get("trace")
	traceContext, _ = trace.(*lib.TraceContext)
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp = Response{
		Code:    code,
		Data:    "",
		TraceId: traceId,
	}
	if err != nil {
		switch c.GetString("lang") {
		case "en":
			resp.Msg = en.getMessage(resp.Code)
		default:
			resp.Msg = zh.getMessage(resp.Code)
		}
	}
	c.JSON(200, resp)
	respone, _ = json.Marshal(resp)
	c.Set("response", string(respone))
	_ = c.AbortWithError(200, err)
	c.Abort()
}

// 返回前端信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	var (
		trace        interface{}
		traceContext *lib.TraceContext
		traceId      string
		response     []byte
		resp         Response
	)
	trace, _ = c.Get("trace")
	traceContext, _ = trace.(*lib.TraceContext)
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp = Response{
		Code:    SuccessCode,
		Msg:     "",
		Data:    data,
		TraceId: traceId,
	}
	switch c.GetString("lang") {
	case "en":
		resp.Msg = en.getMessage(SuccessCode)
	default:
		resp.Msg = zh.getMessage(SuccessCode)
	}
	c.JSON(200, resp)
	response, _ = json.Marshal(resp)
	c.Set("response", string(response))
	c.Next()
}
