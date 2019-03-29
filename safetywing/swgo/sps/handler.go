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
		http.Error(w, "sp id not exist", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	spbytes, err := json.Marshal(sp)
	if err != nil {
		log.Printf(" ** parse spModel to json err: %v\n", err)
	}
	fmt.Fprintf(w, "%s", spbytes)
	return
}

// SearchSps by search key
func SearchSps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	searchKey := ps.ByName("searchKey")
	if searchKey == "" {
		log.Println("  * search sps: empty search key")
		http.Error(w, "empty search key", 400)
		return
	}
	vmap, _ := getSpsBySearchKey(searchKey)
	vbytes, err := json.Marshal(vmap)
	if err != nil {
		log.Println(" ** search sps: vmap to bytes err")
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", vbytes)
	return
}

// UserLastSearchHistory ... user 最近的10个搜索历史
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

// GetAllHis ...
func GetAllHis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	urlVals := r.URL.Query()
	var start string
	start = urlVals.Get("start")
	if start == "" {
		start = "0"
	}
	his := getHisbyStart(start)
	hisbytes, err := json.Marshal(his)
	if err != nil {
		log.Println(" *  get all his parse his to bytes error")
		log.Println(err)
		http.Error(w, "get all his parse his to bytes error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", hisbytes)
	return
}

// GetOneSpHis by spid should have start onece 20 result
func GetOneSpHis(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	spid := ps.ByName("spid")
	urlVals := r.URL.Query()
	var start string
	start = urlVals.Get("start")
	if spid == "" {
		http.Error(w, "GetOneSpHis: empry spid", 400)
		return
	}

	if start == "" {
		start = "0"
	}

	hisVal := getOneSpsHis(spid, start)
	hisJSON, err := json.Marshal(hisVal)
	if err != nil {
		http.Error(w, "GetOneSpHis: parse val to json error", 500)
		log.Printf("GetOneSpHis: parse val to json error, err: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", hisJSON)
	return
}
