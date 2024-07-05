package dao

import "github.com/xiaka53/AdminHub/public"

type UploadSetVisible uint

const (
	UploadSetVisibleNotSelected UploadSetVisible = iota // 未选用
	UploadSetVisibleQiniu                               // 七牛云配置
	UploadSetVisibleAli                                 // 阿里配置
	UploadSetVisibleTx                                  // 腾讯配置
	UploadSetVisibleLoc                                 // 本地配置
)

type Qiniu struct {
	AK     string `form:"AK" json:"AK" validate:"omitempty,min=0" zh:"AK"`
	SK     string `form:"SK" json:"SK" validate:"omitempty,min=0" zh:"SK"`
	Bucket string `form:"bucket" json:"bucket" validate:"omitempty,min=0" zh:"仓库名称"`
	Domain string `form:"domain" json:"domain" validate:"omitempty,min=0" zh:"域名"`
}

type Ali struct {
	Qiniu
	Endpoint string `form:"endpoint" json:"endpoint" validate:"omitempty,min=0" zh:"自定义域名"`
}

type Tx struct {
	Qiniu
	BucketName string `form:"bucketName" json:"bucketName" validate:"omitempty,min=0" zh:"自定义域名"`
}

type UploadSet struct {
	Id      uint             `json:"id" gorm:"column:id;primaryKey;type:int(12) auto_increment;comment:'ID'"`
	Qiniu   string           `gorm:"column:qiniu;type:text;not null;comment:'七牛云'"`
	Alioss  string           `gorm:"column:alioss;type:text;not null;comment:'阿里'"`
	Txcos   string           `gorm:"column:txcos;type:text;not null;comment:'腾讯'"`
	Visible UploadSetVisible `gorm:"column:visible;type:tinyint(1);comment:'配置类型'"`
}

func (u UploadSet) TableName() string {
	return "upload_set"
}

func (u *UploadSet) First() error {
	return public.MainSql.Table(u.TableName()).Where(u).First(u).Error
}

func (u *UploadSet) Edit() error {
	return public.MainSql.Table(u.TableName()).Where("id=?", u.Id).Save(u).Error
}
