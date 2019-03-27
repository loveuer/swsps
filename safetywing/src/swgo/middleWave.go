package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type uMid func(httprouter.Handle) httprouter.Handle

var (
	middleWaves []uMid
	timeBase    = "2006/01/02 15:04:05"
)

func init() {
	middleWaves = append(middleWaves, logger())
}

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
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

type logEntry struct {
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

func logger() uMid {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			start := time.Now()
			sw := statusWriter{ResponseWriter: w}
			next(w, r, ps)
			duration := time.Since(start)
			var statusCode int
			if sw.status == 0 {
				statusCode = 200
			} else {
				statusCode = sw.status
			}
			printLog(logEntry{Host: r.Host, RemoteAddr: r.RemoteAddr, Method: r.Method, RequestURI: r.RequestURI, Proto: r.Proto, Status: statusCode, ContentLen: sw.length, UserAgent: r.Header.Get("User-Agent"), Duration: duration})
			return
		}
	}
}

func printLog(l logEntry) {
	fmt.Printf("%s - - [%s] [%s] %s %s %d -\n", l.Host, time.Now().Format(timeBase), l.Method, l.RequestURI, l.Proto, l.Status)
}

func chain(f httprouter.Handle) httprouter.Handle {
	for _, m := range middleWaves {
		f = m(f)
	}
	return f
}
