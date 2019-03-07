package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"./ops"
	"./session"
	"./sps"
	"./swuser"
	"github.com/julienschmidt/httprouter"
)

var (
	err error
)

func main() {
	r := httprouter.New()

	r.GET("/", chain(home))
	r.GET("/api/session/:sid", chain(sess))
	r.GET("/api/user/:uname", chain(user))
	r.POST("/api/login", chain(login))
	r.GET("/api/history/search/last/:uid", chain(sps.UserLastSearchHistory))
	r.GET("/api/sps/:spid", chain(sps.GetOneSps))

	r.ServeFiles("/static/*filepath", http.Dir("./static"))
	fmt.Println("listening at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", r)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, nil)
}

func sess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sid := ps.ByName("sid")
	if sid == "" || sid == "0" {
		http.Error(w, "Empty session id", 400)
		return
	}

	uname := session.Check(sid)
	if uname == "" {
		http.Error(w, "Wrong session id", 400)
		return
	}
	rjson := &ops.RespJSON{Msg: "200", Val: uname}
	rstr, _ := json.Marshal(rjson)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rstr)
	return
}

// get user by username
func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uname := ps.ByName("uname")
	if uname == "" {
		http.Error(w, "Empty user id", 400)
	}

	u, err := swuser.Get(uname)
	if err != nil {
		http.Error(w, "User id not exist", 400)
		return
	}

	ustr, _ := json.Marshal(u)
	rjson := &ops.RespJSON{Msg: "success", Val: string(ustr)}
	rstr, _ := json.Marshal(rjson)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rstr)
	return
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Wrong request", 400)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Printf("u: %s,  p: %s\n", username, password)
	if username == "" || password == "" {
		http.Error(w, "Wrong args", 400)
		return
	}

	result := swuser.Check(username, password)
	if result {
		w.Header().Set("Content-Type", "application/json")
		rjson := &ops.RespJSON{Msg: "Authenticated", Val: session.New(username)}
		rstr, _ := json.Marshal(rjson)
		fmt.Fprintf(w, "%s", rstr)
		return
	}
	http.Error(w, "Unauthorized", 403)
	return
}
