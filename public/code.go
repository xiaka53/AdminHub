package public

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
)

func GetCodeImage() (string, string) {
	cp := NewCaptcha(152, 61, 4)
	code, img := cp.OutPut()
	base64Str, err := encodeImageToBase64(img)
	if err != nil {
		return "", ""
	}
	return fmt.Sprintf("data:image/jpeg;base64,%v", base64Str), code

}

func encodeImageToBase64(img image.Image) (string, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
