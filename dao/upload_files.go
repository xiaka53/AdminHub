package dao

import (
	"github.com/xiaka53/AdminHub/public"
	"time"
)

type UploadFilesDomain uint
type UploadFilesType uint

const (
	UploadFilesDomainLoc   UploadFilesDomain = iota // 本地配置
	UploadFilesDomainQiniu                          // 七牛云配置
	UploadFilesDomainAli                            // 阿里配置
	UploadFilesDomainTx                             // 腾讯配置
)
const (
	UploadFilesTypePic    UploadFilesType = iota + 1 // 图片
	UploadFilesTypeVidio                             // 视频
	UploadFilesTypeExcel                             // Excel
	UploadFilesTypeWord                              // Word
	UploadFilesTypePdf                               // pdf
	UploadFilesTypeZip                               // zip
	UploadFilesTypeUnNo                              // 未知
	UploadFilesTypeFolder                            // 文件夹
)

type UploadFiles struct {
	Id     uint              `gorm:"column:id;primaryKey;type:int(12) auto_increment;comment:'ID'"`
	Domain UploadFilesDomain `gorm:"column:domain;type:tinyint(1);comment:'保存方式'"`
	Type   UploadFilesType   `gorm:"column:type;type:tinyint;comment:'类型'"`
	Name   string            `gorm:"column:name;type:varchar(255);comment:'文件名称'"`
	Key    string            `gorm:"column:key;type:varchar(255);comment:'三方的Key'"`
	Url    string            `gorm:"column:url;type:varchar(255);comment:'文件地址'"`
	Atime  time.Time         `gorm:"column:atime;type:datetime;comment:'创建时间'"`
	Pid    uint              `gorm:"column:pid;type:tinyint;comment:'上级目录ID'"`
}

func (u UploadFiles) TableName() string {
	return "upload_files"
}

func (u *UploadFiles) First() error {
	return public.MainSql.Table(u.TableName()).Where(u).First(u).Error
}

func (u *UploadFiles) Total() (total int) {
	public.MainSql.Table(u.TableName()).Where(u).Count(&total)
	return total
}

func (u *UploadFiles) FromPage(page, size int, order, name string) (data []UploadFiles, total int) {
	if order == "" {
		order = "id asc"
	}
	where := "1=1"
	args := []any{}
	if name != "" {
		where = where + " and name like '%?%'"
		args = append(args, name)
	}
	if u.Pid == 0 {
		where = where + " and `pid` =0"
	}
	if u.Type > 0 {
		where = where + " and `type` in (?,8)"
		args = append(args, u.Type)
		u.Type = 0
	}
	subQuery := public.MainSql.Table(u.TableName()).Where(u).Where(where, args...)
	subQuery.Limit(size).Offset(size * (page - 1)).Order(order).Find(&data)
	subQuery.Count(&total)
	return
}

func (u *UploadFiles) Create() error {
	u.Atime = time.Now()
	return public.MainSql.Table(u.TableName()).Create(u).Error
}

func (u *UploadFiles) Edit() error {
	return public.MainSql.Table(u.TableName()).Where("id=?", u.Id).Updates(u).Error
}

func (u *UploadFiles) Del() error {
	return public.MainSql.Table(u.TableName()).Delete(u).Error
}
