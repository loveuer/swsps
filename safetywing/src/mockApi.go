package main

import (
	"encoding/json"
	"fmt"
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
}

func main() {
	m := mux.NewRouter()

	m.HandleFunc("/api/session/{id}", chain(sess)).Methods("GET")
	m.HandleFunc("/api/session/", chain(sess)).Methods("GET")

	fmt.Println("listening at http://localhost:8000")
	http.ListenAndServe("127.0.0.1:8000", m)
}

func sess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionid := vars["id"]

	var rs []byte
	if sessionid == "" {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: ""})
	} else {
		rs, _ = json.Marshal(RespJSON{StatusCode: 200, Msg: "loveuer"})
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", rs)
	return
}
