package sps

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../ops"
	"github.com/julienschmidt/httprouter"
)

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
