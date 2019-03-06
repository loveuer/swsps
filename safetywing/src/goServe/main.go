package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"./session"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", chain(home))
	r.GET("/api/session/:sid", chain(sess))
	r.GET("/api/user/:uid", chain(user))

	fmt.Println("listening at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", r)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, nil)
}

func sess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sid := ps.ByName("sid")
	if sid == "" {
		http.Error(w, "empty sid", 400)
		return
	}

	uid := session.Check(sid)
	if uid == 0 {
		http.Error(w, "wrong sid", 400)
		return
	}
	rjson := &RespJSON{Msg: "200", Val: strconv.Itoa(uid)}
	rstr, _ := json.Marshal(rjson)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rstr)
	return
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if uid == "" {
		http.Error(w, "empty uid", 400)
	}

	u, ok := user.Get(uid)
}
