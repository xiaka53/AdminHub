package public

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/DeployAndLog/lib"
)

// 普通日志
func ComLogNotice(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagInfo(traceContext, dltag, m)
}

// 普通日志
func ComLogError(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagError(traceContext, dltag, m)
}

// 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *lib.TraceContext {
	// 防御
	if c == nil {
		return lib.NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*lib.TraceContext); ok {
			return tc
		}
	}
	return lib.NewTrace()
}

// 获取新的TraceId载体Context
func GetNewTraceContext(TraceId string) *gin.Context {
	var c gin.Context
	traceContext := lib.NewTrace()
	traceContext.TraceId = TraceId
	c.Set("trace", traceContext)
	return &c
}
