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
	val    int
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
func New() string {
	rand.Seed(time.Now().UnixNano())
	sessKey := uuid()
	sessMap.Store(sessKey, time.Now().Add(1*time.Hour))
	return sessKey
}

// Check session id if in map and if expired
func Check(sid string) int {
	val, ok := sessMap.Load(sid)
	if !ok {
		return 0
	}
	if val.(sessVal).expire.Before(time.Now()) {
		sessMap.Delete(sid)
		return 0
	}
	return val.(sessVal).val
}
