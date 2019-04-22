package salary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"../socket"

	"github.com/julienschmidt/httprouter"
)

// Upload <- salary
func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err = r.ParseMultipartForm(16 << 20)
	if err != nil {
		http.Error(w, "salary -> Upload: ParseMultipartForm Error", 400)
		return
	}

	file, fHandler, err := r.FormFile("xlsx")
	if err != nil {
		http.Error(w, "salary -> Upload: FormFile Error", 400)
		return
	}

	if len(strings.Split(fHandler.Filename, ".")) != 2 {
		// file name err
		http.Error(w, "salary -> Upload: filename split != 2 Error", 400)
		return
	}

	fileName := strings.Split(fHandler.Filename, ".")[0]
	fileType := strings.Split(fHandler.Filename, ".")[1]

	if fileType != "xlsx" {
		// file type err
		http.Error(w, "salary -> Upload: file type != xlsx Error", 400)
		return
	}

	ok, _ := regexp.Match(`20\d{2}年(0[1-9]|1[0-2])月工资条`, []byte(fileName))
	if !ok {
		// file type err / not match regex
		http.Error(w, "salary -> Upload: filename not match regex Error", 400)
		return
	}

	filePath := strings.Split(fileName, "工资条")[0]
	folderPath := path.Join("/home/codes/golang/uhome/salary/xlsx", filePath)
	err = os.MkdirAll(folderPath, 0777)
	if err != nil {
		// make upload folder err
		http.Error(w, "salary -> Upload: make new folder Error", 400)
		return
	}

	newFile, err := os.OpenFile(folderPath+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		// make(open) new file err
		http.Error(w, "salary -> Upload: open new file Error", 400)
		return
	}

	io.Copy(newFile, file)

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"event":  "salary",
		"type":   "upload",
		"status": "success",
	})
	fmt.Fprintf(w, "%s", resp)
	return
}

func RunScript(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "salary -> RunScript: parse form err", 400)
		return
	}

	yearmonth := r.Form.Get("yearmonth")
	line := r.Form.Get("line")
	if yearmonth == "" {
		http.Error(w, "salary -> RunScript: year and month are empty", 400)
		return
	}
	if line == "" {
		line = "2"
	}

	ok, _ := regexp.Match(`20\d{2}(0[1-9]|1[0-2])`, []byte(yearmonth))
	if !ok {
		http.Error(w, "salary -> RunScript: argv(year and month not match regex)", 400)
		return
	}

	folderName := yearmonth[0:4] + "年" + yearmonth[4:6] + "月"
	cmd := exec.Command("./salary/xlsx/testCmd.py", folderName, line)
	err = cmd.Start()
	if err != nil {
		http.Error(w, "salary -> RunScript: cmd start err", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"event":      "salary",
		"type":       "run script",
		"status":     "success",
		"process_id": strconv.Itoa(cmd.Process.Pid),
	})
	fmt.Fprintf(w, "%s", resp)
	return
}

func CheckScript(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "salary -> CheckScript: parse form err", 400)
		return
	}

	yearmonth := r.Form.Get("yearmonth")
	if yearmonth == "" {
		http.Error(w, "salary -> CheckScript: year and month are empty", 400)
		return
	}

	ok, _ := regexp.Match(`20\d{2}(0[1-9]|1[0-2])`, []byte(yearmonth))
	if !ok {
		http.Error(w, "salary -> CheckScript: argv(year and month not match regex)", 400)
		return
	}

	go sendRuningResult(yearmonth)

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"event":  "salary",
		"type":   "check script",
		"status": "success",
	})
	fmt.Fprintf(w, "%s", resp)
	return
}

func sendRuningResult(ym string) {
	for {
		if !reading {
			socket.UhomeSocket.Sender <- "close reading"
			break
		}
		resp, _ := json.Marshal(map[string]string{
			"event": "salary",
			"type":  "update check log",
			"data":  fmt.Sprintf("%s: line xx done", ym),
		})
		socket.UhomeSocket.Sender <- string(resp)
		time.Sleep(5 * time.Second)
	}
}
