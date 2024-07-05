package upload

import (
	"github.com/xiaka53/AdminHub/public"
	"github.com/xiaka53/AdminHub/redis"
	"os"
)

type Local struct {
}

func GetLoc() *Local {
	return &Local{}
}

func (l *Local) GetToken() map[string]any {
	data := make(map[string]any)
	uuid := public.GetUUid()
	data["uptoken"] = uuid
	(&redis.Upload{
		Uuid:   uuid,
		Expire: 3600,
	}).SetUuid()
	return data
}

func (l *Local) Delete(name string) {
	_ = os.Remove("files/" + name)
}
