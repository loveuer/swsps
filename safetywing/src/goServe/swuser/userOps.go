package swuser

import (
	"time"

	"../ops"
	_ "github.com/lib/pq"
)

// User user base type
type User struct {
	ID          int       `json:"id"`
	UserName    string    `json:"username"`
	RealName    string    `json:"realname"`
	ProfileIcon string    `json:"profile_icon"`
	LastLogin   time.Time `json:"last_login"`
}

var (
	conn        = ops.Conn
	dbConfig    = "host=192.168.31.99 port=8765 user=loveuer password=1314loveuer dbname=loveuer sslmode=require"
	err         error
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)
