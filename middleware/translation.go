package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
)

// 设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("lang")
		if locale == "" {
			locale = "zh"
		}
		trans, _ := public.Uni.GetTranslator(locale)
		// 设置trans到context
		c.Set("trans", trans)
		c.Set("lang", locale)
		c.Next()
	}
}
