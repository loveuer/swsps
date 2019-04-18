package main

import (
	"fmt"
	"net/http"

	"./salary"
	"./socket"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.POST("/api/salary/upload", salary.Upload)
	r.GET("/api/socket", socket.Init)

	r.ServeFiles("/static/*filepath", http.Dir("static"))
	fmt.Println("listening at http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
