package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/xiaka53/DeployAndLog/lib"
)

func getServer() redis.Conn {
	rs, err := lib.RedisConnFactory("base")
	if err != nil {
		return nil
	}
	return rs
}

func delConn(rs redis.Conn) {
	if rs == nil {
		return
	}
	_ = rs.Close()
	return
}
