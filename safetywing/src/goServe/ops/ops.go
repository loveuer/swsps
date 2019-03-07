package ops

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// RespJSON normal response
type RespJSON struct {
	Msg string `json:"msg"`
	Val string `json:"val"`
}

var (
	// Conn to postgres db
	Conn     *sql.DB
	dbConfig = "host=192.168.31.99 port=8765 user=loveuer password=1314loveuer dbname=loveuer sslmode=require"
	err      error
)

func init() {
	Conn, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatal(" *** can't open db")
	}
}
