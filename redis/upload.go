package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Upload struct {
	Uuid   string
	Expire int
}

func (u Upload) DB() redis.Conn {
	rs := getServer()
	if rs == nil {
		return nil
	}
	if _, err := rs.Do("select", 1); err != nil {
		return nil
	}
	return rs
}

func (u Upload) SetUuid() {
	key := fmt.Sprintf("upload_uuid:%s", u.Uuid)
	rs := u.DB()
	defer delConn(rs)
	if _, err := rs.Do("set", key, 1); err != nil {
		return
	}
	_, _ = rs.Do("expire", key, u.Expire)
}

func (u *Upload) GetUuid() bool {
	key := fmt.Sprintf("upload_uuid:%s", u.Uuid)
	rs := u.DB()
	defer delConn(rs)
	code, err := redis.Int(rs.Do("get", key))
	if err != nil {
		return false
	}
	if code != 1 {
		return false
	}
	return true
}
