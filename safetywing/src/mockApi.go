package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

// LogEntry ...
type LogEntry struct {
	Host       string
	RemoteAddr string
	Method     string
	RequestURI string
	Proto      string
	Status     int
	ContentLen int
	UserAgent  string
	Duration   time.Duration
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

type uMid func(http.HandlerFunc) http.HandlerFunc

var (
	mids     []uMid
	timeBase = "2006/01/02 15:04:05"
)

func chain(f http.HandlerFunc) http.HandlerFunc {
	for _, m := range mids {
		f = m(f)
	}

	return f
}

func logger() uMid {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sw := statusWriter{ResponseWriter: w}
			duration := time.Since(start)
			f(w, r)
			printLog(LogEntry{
				Host:       r.Host,
				RemoteAddr: r.RemoteAddr,
				Method:     r.Method,
				RequestURI: r.RequestURI,
				Proto:      r.Proto,
				Status:     sw.status,
				ContentLen: sw.length,
				UserAgent:  r.Header.Get("User-Agent"),
				Duration:   duration,
			})
			return
		}
	}
}

func printLog(l LogEntry) {
	fmt.Printf("%s - - [%s] [%s] %s %s %d -\n", l.Host, time.Now().Format(timeBase), l.Method, l.RequestURI, l.Proto, l.Status)
}

func init() {
	mids = append(mids, logger())
}

// RespJSON ...
type RespJSON struct {
	StatusCode int
	Msg        string
	Val        string
}

func main() {
	m := mux.NewRouter()

	// handle session chk if logined
	m.HandleFunc("/api/session/{id}", chain(sess)).Methods("GET")
	m.HandleFunc("/api/session/", chain(sess)).Methods("GET")
	// post to login user
	m.HandleFunc("/api/login", chain(login)).Methods("POST")
	// after logined get user base imf
	m.HandleFunc("/api/user/{id}", chain(user)).Methods("GET")
	// user search history
	m.HandleFunc("/api/history/most/{id}", chain(historyMost)).Methods("GET")

	m.PathPrefix("/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Println("listening at http://localhost:8000")
	http.ListenAndServe("127.0.0.1:8000", m)
}

func sess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionid := vars["id"]

	var rs []byte
	if sessionid == "new session" {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "valid session", Val: "1"})
	} else {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "wrong session", Val: ""})
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rs)
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Username string `json: "username"`
		Password string `json: "password"`
		Csrf     string `json: "csrf"`
	}
	var (
		err error
		p   Params
		rs  []byte
	)

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &p)

	fmt.Printf("username: %s, password: %s\n", p.Username, p.Password)

	if err != nil {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "wrong params"})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", rs)
		return
	}

	if p.Username == "loveuer" && p.Password == "1314loveuer" {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "200", Val: "new session"})
	} else {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "auth fail"})
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rs)
	return
}

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Uuser no...
	type Uuser struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Realname string `json:"realname"`
		Img      string `json:"img"`
	}

	var (
		rs []byte
	)
	w.Header().Set("Content-Type", "application/json")
	if id == "0" || id == "" {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "wrong id", Val: ""})
		fmt.Fprintf(w, "%s", rs)
		return
	}
	if id == "1" {
		u, _ := json.Marshal(Uuser{ID: 1, Username: "loveuer", Realname: "赵育鹏", Img: "http://localhost:8000/static/头像.jpg"})
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "200", Val: string(u)})
		fmt.Fprintf(w, "%s", rs)
	} else {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "wrong id", Val: ""})
		fmt.Fprintf(w, "%s", rs)
	}
	return
}

func historyMost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "wrong user id", 400)
		return
	}

	mockResp := []string{
		"pdf", "mcdu", "du", "5v", "test",
	}
	w.Header().Set("Content-Type", "application/json")
	rp, _ := json.Marshal(mockResp)
	fmt.Fprintf(w, "%s", rp)
	return
}
