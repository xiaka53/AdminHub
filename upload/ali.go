package upload

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"
)

type Ali struct {
	AK       string
	SK       string
	Bucket   string
	Domain   string
	Endpoint string
	Expires  time.Duration
}

func GetAli(_ali string) *Ali {
	var base Ali
	_ = json.Unmarshal([]byte(_ali), &base)
	base.Expires = time.Hour + 1
	return &base
}

type Policy struct {
	Expiration string          `json:"expiration"`
	Conditions [][]interface{} `json:"conditions"`
}

func (a *Ali) GetToken() map[string]any {
	data := make(map[string]any)
	expiration := time.Now().UTC().Add(30 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	conditions := [][]interface{}{
		{"content-length-range", 0, 1048576000},
	}
	policy := Policy{
		Expiration: expiration,
		Conditions: conditions,
	}
	policyBytes, err := json.Marshal(policy)
	if err != nil {
		return data
	}
	policyBase64 := base64.StdEncoding.EncodeToString(policyBytes)
	h := hmac.New(sha1.New, []byte(a.SK))
	h.Write([]byte(policyBase64))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	data = map[string]any{
		"policy":         policyBase64,
		"signature":      signature,
		"OSSAccessKeyId": a.AK,
		"domain":         a.Domain,
	}
	return data
}

func (a *Ali) Delete(name string) {
	client, err := oss.New(a.Endpoint, a.AK, a.SK)
	if err != nil {
		return
	}
	bucket, err := client.Bucket(a.Bucket)
	if err != nil {
		return
	}
	err = bucket.DeleteObject(name)
	if err != nil {
		return
	}
	return
}
