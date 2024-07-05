package templae

const TemplateDaoStr = `package dao

import (
	"github.com/xiaka53/AdminHub/public"
)

type {{ .StructName }} struct {
{{- range .Columns }}
	{{ .ColumnNameType }} {{ .GoType }} ` + "`json:\"{{ .ColumnName }}\" gorm:\"column:{{ .ColumnName }}\"`" + ` // {{ .ColumnComment }}
{{- end }}
}

func ({{ .ShortName }} *{{ .StructName }}) TableName() string {
	return "{{ .TableName }}"
}

func ({{ .ShortName }} *{{ .StructName }}) First() error {
	return public.MainSql.Table({{ .ShortName }}.TableName()).Where({{ .ShortName }}).First({{ .ShortName }}).Error
}

func ({{ .ShortName }} *{{ .StructName }}) FromPage(page, size int, order string) (data []{{ .StructName }}, total int) {
	if order == "" {
		order = "id asc"
	}
	public.MainSql.Table({{ .ShortName }}.TableName()).Where({{ .ShortName }}).Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	public.MainSql.Table({{ .ShortName }}.TableName()).Where({{ .ShortName }}).Count(&total)
	return
}

func ({{ .ShortName }} *{{ .StructName }}) Create() error {
	return public.MainSql.Table({{ .ShortName }}.TableName()).Create({{ .ShortName }}).Error
}

func ({{ .ShortName }} *{{ .StructName }}) Edit() error {
	return public.MainSql.Table({{ .ShortName }}.TableName()).Where("id=?", {{ .ShortName }}.Id).Updates({{ .ShortName }}).Error
}

func ({{ .ShortName }} *{{ .StructName }}) Del() error {
	return public.MainSql.Table({{ .ShortName }}.TableName()).Delete({{ .ShortName }}).Error
}
`
