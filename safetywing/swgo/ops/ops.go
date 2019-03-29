package ops

import (
	"database/sql"
	"fmt"
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
	dbConfig = "host=192.168.31.99 port=8765 user=loveuer password=1314loveuer dbname=swdb sslmode=require"
	// dbConfig = "host=118.113.72.134 port=8765 user=loveuer password=1314loveuer dbname=swdb sslmode=require"
	err error
)

func init() {
	fmt.Println("请输出database ip( defaul: 192.168.31.99 ):")
	var ipinput string
	fmt.Scanln(&ipinput)
	if ipinput != "" {
		dbConfig = fmt.Sprintf("host=%s port=8765 user=loveuer password=1314loveuer dbname=swdb sslmode=require", ipinput)
	}
	Conn, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatal(" *** can't open db")
	}
}
