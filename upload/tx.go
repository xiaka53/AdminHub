package upload

import (
	"context"
	"encoding/json"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"net/http"
	"net/url"
	"time"
)

type Tx struct {
	AK         string
	SK         string
	Bucket     string
	Domain     string
	BucketName string
	Expires    time.Duration
}

func GetTx(_tx string) *Tx {
	var base Tx
	_ = json.Unmarshal([]byte(_tx), &base)
	base.Expires = time.Hour + 1
	return &base
}

func (t *Tx) GetToken() map[string]any {
	data := make(map[string]any)
	client := sts.NewClient(t.AK, t.SK, nil)
	opt := &sts.CredentialOptions{
		DurationSeconds: 1800,
		Region:          t.Bucket,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + t.Bucket + ":uid/" + "1252001357" + ":" + t.BucketName + "/",
					},
				},
			},
		},
	}
	res, err := client.GetCredential(opt)
	if err != nil {
		return data
	}

	data = map[string]any{
		"Bucket":        t.BucketName,
		"ExpiredTime":   res.ExpiredTime,
		"Region":        t.Bucket,
		"SecurityToken": res.Credentials.SessionToken,
		"StartTime":     res.StartTime,
		"TmpSecretId":   res.Credentials.TmpSecretID,
		"TmpSecretKey":  res.Credentials.TmpSecretKey,
		"path":          t.Domain,
	}
	return data
}

func (t *Tx) Delete(name string) {
	u, _ := url.Parse(t.Domain)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.AK,
			SecretKey: t.SK,
		},
	})
	_, _ = client.Object.Delete(context.Background(), name)
	return
}
