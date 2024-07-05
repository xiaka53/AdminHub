package redis

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	SigningKey = "34rf324gredsniwer01"
)

type Login struct {
	Uuid    string
	AdminId uint
	Code    string
	Token   string
}

func (l Login) DB() redis.Conn {
	rs := getServer()
	if rs == nil {
		return nil
	}
	if _, err := rs.Do("select", 0); err != nil {
		return nil
	}
	return rs
}

func (l Login) SetUuid() {
	key := fmt.Sprintf("uuid:%s", l.Uuid)
	rs := l.DB()
	defer delConn(rs)
	if _, err := rs.Do("set", key, l.Code); err != nil {
		return
	}
	_, _ = rs.Do("expire", key, 5*60)
}

func (l *Login) GetCode() {
	key := fmt.Sprintf("uuid:%s", l.Uuid)
	rs := l.DB()
	defer delConn(rs)
	code, err := redis.String(rs.Do("get", key))
	if err != nil {
		return
	}
	l.Code = code
	_, _ = rs.Do("del", key)
}

func (l *Login) SetToken() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id":    l.AdminId,
		"expire_time": time.Now().Add(3600 * 24 * 30 * time.Second).Unix(),
	})
	var err error
	l.Token, err = token.SignedString([]byte(SigningKey))
	if err != nil {
		return
	}
	key := fmt.Sprintf("admin_login_token:%d", l.AdminId)
	rs := l.DB()
	defer delConn(rs)
	if _, err = rs.Do("set", key, l.Token); err != nil {
		return
	}
	_, _ = rs.Do("expire", key, 3600*24*30)
}

func (l *Login) GetToken() {
	token, err := jwt.Parse(l.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SigningKey), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if expire_time, _ok := claims["expire_time"]; _ok {
			if int64(expire_time.(float64)) < time.Now().Unix() {
				return
			}
		}
		if adminId, _ok := claims["admin_id"]; _ok {
			l.AdminId = uint(adminId.(float64))
		}
	}
	if l.AdminId < 1 {
		return
	}
	key := fmt.Sprintf("admin_login_token:%d", l.AdminId)
	rs := l.DB()
	defer delConn(rs)
	_token, err := redis.String(rs.Do("get", key))
	if err != nil || _token != l.Token {
		l.AdminId = 0
		return
	}
	return
}
