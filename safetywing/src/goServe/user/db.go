package user

import (
	"database/sql"
	"log"
	"time"

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
	conn        *sql.DB
	dbConfig    = "host=192.168.31.99 port=8765 user=loveuer password=1314loveuer dbname=loveuer sslmode=require"
	err         error
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	conn, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatal("can't open db")
	}
}
