package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"./ops"
	"./session"

	"./socket"
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
	// about user
	r.GET("/api/history/search/last/:uid", chain(sps.UserLastSearchHistory))
	// aboyt sps
	r.GET("/api/sps/one/:spid", chain(sps.GetOneSps))
	r.GET("/api/sps/search/:searchKey", chain(sps.SearchSps))
	// about sps his
	r.GET("/api/sphis", chain(sps.GetAllHis))
	r.GET("/api/sphis/one/:spid", chain(sps.GetOneSpHis))
	// about log
	r.POST("/api/log/add", chain(NewLog))

	// about websocket
	go socket.Manager.Start()
	r.GET("/api/ws", chain(socket.Upgrade2WS))
	r.POST("/api/ws/mockPost", chain(socket.MockPostSome))
	// r.Handler()

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
	rjson := &ops.RespJSON{Msg: "accepted", Val: uname}
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

	ubytes, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", ubytes)
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
		ruser, _ := swuser.Get(username)
		rsess := session.New(username)
		rmap := map[string]interface{}{"session": rsess, "user": ruser}
		rbytes, err := json.Marshal(rmap)
		if err != nil {
			log.Println(" *** login: rmap to rbytes error")
			log.Println(err)
			http.Error(w, "login: rmap to rbytes error", 500)
			return
		}
		fmt.Fprintf(w, "%s", rbytes)
		return
	}
	http.Error(w, "Unauthorized", 403)
	return
}

// NewLog ...
// test 一下 websocket
func NewLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err = r.ParseForm()
	if err != nil {
		log.Printf("(main NewLog) parse form err: %v\n", err)
		http.Error(w, "parse form error, wrong params", 400)
		return
	}

	sim := r.Form.Get("sim")
	log := r.Form.Get("log")
	wsid := r.Form.Get("wsid")

	fmt.Printf("sim: %s\nlog: %s\nid: %s\n", sim, log, wsid)

	if wsid == "" {
		http.Error(w, "未认证的 ws id", 400)
		return
	}
	if sim == "" || log == "" {
		http.Error(w, "sim or log empty", 400)
	}

	// add log to db

	// send add message by wbsocket
	wbMess, _ := json.Marshal(map[string]string{"sim": sim, "log": log})
	wbResp := map[string]string{
		"id":      wsid,
		"type":    "newlog",
		"message": string(wbMess),
	}
	socket.Manager.Send(wbResp)
	return
}
