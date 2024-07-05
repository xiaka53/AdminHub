package templae

const TemplateApiStr = `package api

import (
	"{{ .ServerName }}/dao"
	"{{ .ServerName }}/dto"
	"github.com/xiaka53/AdminHub/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type {{ .StructName }} struct {
}

func {{ .StructTableName }}RouterGroup(r *gin.RouterGroup) {
	var s {{ .StructName }}
	r.POST("{{ .TableName }}/list", s.{{ .StructName }}List)
	r.POST("{{ .TableName }}/add", s.{{ .StructName }}Add)
	r.POST("{{ .TableName }}/edit", s.{{ .StructName }}Edit)
	r.POST("{{ .TableName }}/del", s.{{ .StructName }}Del)
}

func ({{ .StructName }}) {{ .StructName }}List(c *gin.Context) {
	var (
		params dto.{{ .StructTableName }}List
		info   {{ .StructName }}ListInfo
		data   []dao.{{ .StructTableName }}
		err    error
	)
	if err = (&params).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	data, info.All = (&dao.{{ .StructTableName }}{}).FromPage(params.Page.Page, params.Size, params.OrderBy.OrderBy)
	for _, v := range data {
		info.Datas = append(info.Datas, dao.{{ .StructTableName }}{
			{{ .PrimaryKey }}:v.{{ .PrimaryKey }},
		{{- range .Columns }}
			{{ .ColumnNameType }}: v.{{ .ColumnNameType }},
		{{- end }}
		})
	}
	middleware.ResponseSuccess(c, info)
}

type {{ .StructName }}ListInfo struct {
	All   int                ` + "`json:\"all\"`" + `
	Datas []dao.{{ .StructTableName}} ` + "`json:\"datas\"`" + `
}


func ({{ .StructName }}) {{ .StructName }}Add(c *gin.Context) {
	var (
		params dto.{{ .StructTableName }}Add
		{{ .LowerCamelTableName }} dao.{{ .StructTableName }}
		err    error
	)
	if err = (&params).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	{{ .LowerCamelTableName }} = dao.{{ .StructTableName }}{
		{{- range .Columns }}
			{{ .ColumnNameType }}: params.{{ .ColumnNameType }},
		{{- end }}
	}
	if err = (&{{ .LowerCamelTableName }}).Create(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func ({{ .StructName }}) {{ .StructName }}Edit(c *gin.Context) {
	var (
		params dto.{{ .StructTableName }}Edit
		{{ .LowerCamelTableName }} dao.{{ .StructTableName }}
		err    error
	)
	if err = (&params).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	{{ .LowerCamelTableName }}.{{ .PrimaryKey }} = params.{{ .PrimaryKey }}
	if err = (&{{ .LowerCamelTableName }}).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
{{- range .Columns }}
	{{ .LowerCamelTableName }}.{{ .ColumnNameType }}= params.{{ .ColumnNameType }}
{{- end }}
	if err = (&{{ .LowerCamelTableName }}).Edit(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}

func ({{ .StructName }}) {{ .StructName }}Del(c *gin.Context) {
	var (
		params dto.{{ .StructTableName }}Del
		{{ .LowerCamelTableName }} dao.{{ .StructTableName }}
		err    error
	)
	if err = (&params).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.InternalErrorCode, err)
		return
	}
	{{ .LowerCamelTableName }}.{{ .PrimaryKey }} = params.{{ .PrimaryKey }}
	if err = (&{{ .LowerCamelTableName }}).First(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	if err = (&{{ .LowerCamelTableName }}).Del(); err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	middleware.ResponseSuccess(c, "")
}
`
