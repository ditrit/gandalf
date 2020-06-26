package upload

import (
	"connectors/gandalf-core/connectors/goworkflow/workflow"
	"crypto/md5"
	"fmt"
	"io"

	goclient "github.com/ditrit/gandalf/libraries/goclient"

	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type UploadController struct {
	url           string
	clientGandalf *goclient.ClientGandalf
}

func NewUploadController(url string, clientGandalf *goclient.ClientGandalf) *UploadController {
	uploadController := new(UploadController)
	uploadController.url = url
	uploadController.clientGandalf = clientGandalf

	return uploadController
}

func (uc UploadController) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		//token := fmt.Sprintf("%x", h.Sum(nil))

		//t, _ := template.ParseFiles("UI/upload.tpl")
		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("upload/tmpl/layout.tmpl", "upload/tmpl/content.tmpl"))
		//tmpl.Execute(w, token)
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Upload workflow", "Url": uc.url})

	} else {
		r.ParseMultipartForm(32 << 20)
		//SOURCE
		sourceFile, sourceHandler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer sourceFile.Close()
		//fmt.Fprintf(w, "%v", sourceHandler.Header)
		//CONFIGURATION
		confFile, confHandler, err := r.FormFile("uploadconf")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer confFile.Close()
		//fmt.Fprintf(w, "%v", confHandler.Header)

		filepath := "./workflows/"
		currentFilePath := filepath + time.Now().Format("01-02-2006 15:04:05") + "/"
		//strconv.FormatInt(time.Now().Unix(), 10)
		os.Mkdir(currentFilePath, 0700)
		fileName := sourceHandler.Filename
		confName := confHandler.Filename

		//SOURCE COPY
		source, err := os.OpenFile(currentFilePath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer source.Close()
		io.Copy(source, sourceFile)

		//CONFIGURATION COPY
		conf, err := os.OpenFile(currentFilePath+confName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conf.Close()
		io.Copy(conf, confFile)

		//ExecuteWorkflow
		workflow.ExecuteWorkflow(currentFilePath, fileName)

		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("upload/tmpl/layout.tmpl", "upload/tmpl/succes.tmpl"))
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Upload succes"})
	}
}
