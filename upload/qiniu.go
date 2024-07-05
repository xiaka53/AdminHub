package upload

import (
	"encoding/json"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiu struct {
	AK      string
	SK      string
	Bucket  string
	Domain  string
	Expires time.Duration
}

func GetQiNiu(_qiniu string) *QiNiu {
	var base QiNiu
	_ = json.Unmarshal([]byte(_qiniu), &base)
	base.Expires = time.Hour + 1
	return &base
}

func (q *QiNiu) GetToken() map[string]any {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	putPolicy.Expires = uint64(time.Now().Add(q.Expires).Unix())
	mac := qbox.NewMac(q.AK, q.SK)
	data := make(map[string]any)
	data["domain"] = q.Domain
	data["token"] = putPolicy.UploadToken(mac)
	return data
}

func (q *QiNiu) Delete(name string) {
	mac := qbox.NewMac(q.AK, q.SK)
	cfg := storage.Config{UseHTTPS: false}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	_ = bucketManager.Delete(q.Bucket, name)
	return
}
