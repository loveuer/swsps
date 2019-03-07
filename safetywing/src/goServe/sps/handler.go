package sps

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../ops"
	"github.com/julienschmidt/httprouter"
)

// GetOneSps get one sp by its spid
func GetOneSps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	spid := ps.ByName("spid")
	if spid == "" {
		http.Error(w, "Empty spid", 400)
		return
	}

	sp, err := getSpByID(spid)
	if err != nil {
		http.Error(w, "spid not exist", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	spStr, err := json.Marshal(sp)
	rjson := &ops.RespJSON{Msg: "success", Val: string(spStr)}
	rstr, err := json.Marshal(rjson)
	if err != nil {
		log.Printf("  * %v\n", err)
	}
	fmt.Fprintf(w, "%s", rstr)
	return
}

// UserLastSearchHistory ... user 最近的5个搜索历史
func UserLastSearchHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if uid == "" {
		http.Error(w, "Empty user id", 400)
		return
	}

	lastHiss, err := getUserLastSearchHistory(uid)
	if err != nil || lastHiss == nil {
		http.Error(w, "Wrong user id or No search history", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	lastHissStr, _ := json.Marshal(lastHiss)
	rjson := &ops.RespJSON{Msg: "success", Val: string(lastHissStr)}
	rstr, _ := json.Marshal(rjson)
	fmt.Fprintf(w, "%s", rstr)
}
