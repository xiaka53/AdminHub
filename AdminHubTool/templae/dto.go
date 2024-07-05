package templae

const TemplateDtoStr = `package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/public"
	"github.com/xiaka53/AdminHub/exec/dto"
	"reflect"
)

type {{ .StructTableName }}List struct {
	dto.OrderBy
	dto.Page
}

func (o *{{ .StructTableName }}List) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type {{ .StructTableName }}Add struct {
	{{- range .Columns }}
		{{ .ColumnNameType }} {{ .GoType }} ` + "`json:\"{{ .ColumnName }}\" form:\"{{ .ColumnName }}\" validate:\"min=1\" zh:\"\" `" + `
	{{- end }}
}

func (o *{{ .StructTableName }}Add) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type {{ .StructTableName }}Edit struct {
	{{ .StructTableName }}Add
	{{ .StructTableName }}Del
}

func (o *{{ .StructTableName }}Edit) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}

type {{ .StructTableName }}Del struct {
	{{ .PrimaryKey }} {{ .PrimaryKeyType }} ` + "`form:\"{{ .PrimaryKeySmall }}\" json:\"{{ .PrimaryKeySmall }}\" validate:\"min=1\" zh:\"{{ .PrimaryKeyName }}\"`" + `
}

func (o *{{ .StructTableName }}Del) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	return public.BindingValidParams(c, o, reflect.TypeOf(*o))
}
`
