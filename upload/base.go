package upload

type UploadFile interface {
	GetToken() map[string]any
	Delete(name string)
}
