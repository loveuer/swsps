package session

import (
	"math/rand"
	"sync"
	"time"
)

var (
	sessMap     sync.Map
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type sessVal struct {
	val    string // username
	expire time.Time
}

func uuid() string {
	rdstr := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}(32)

	return rdstr
}

// New generate a new session
func New(username string) string {
	rand.Seed(time.Now().UnixNano())
	sessKey := uuid()
	sessMap.Store(sessKey, sessVal{val: username, expire: time.Now().Add(1 * time.Hour)})
	return sessKey
}

// Check session id if in map and if expired
func Check(sid string) string {
	val, ok := sessMap.Load(sid)
	if !ok {
		return ""
	}
	if val.(sessVal).expire.Before(time.Now()) {
		sessMap.Delete(sid)
		return ""
	}
	return val.(sessVal).val
}
