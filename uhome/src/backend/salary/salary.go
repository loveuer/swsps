package salary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

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
