package public

import (
	"math/rand"
	"time"
)

const (
	CAPITAL randType = iota
	LOWER_CASE
	NUMBER
	SPECIAL
)

type randType int

type randstruct struct {
	Rand    *rand.Rand
	Str     string
	StrLow  string
	Number  string
	Special string
}

var R randstruct

func init() {
	R = randstruct{
		Rand:    rand.New(rand.NewSource(time.Now().Unix())),
		Str:     "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		StrLow:  "abcdefghijklmnopqrstuvwxyz",
		Number:  "012456789",
		Special: "_[]-<>%!@#$?.。,",
	}
}

//获取随机字符串
func RandString(lens int, a ...randType) (str string) {
	bytes := make([]byte, lens)
	var seed string
	for _, v := range a {
		switch v {
		case CAPITAL:
			seed += R.Str
		case LOWER_CASE:
			seed += R.StrLow
		case NUMBER:
			seed += R.Number
		case SPECIAL:
			seed += R.Special
		}
	}
	seedBytes := []byte(seed)
	for i := 0; i < lens; i++ {
		bytes[i] = seedBytes[R.Rand.Intn(len(seed)-1)]
	}
	str = string(bytes)
	return
}
