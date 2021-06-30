package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/connectors/goworkflowdocker/docker"
	"github.com/ditrit/gandalf/connectors/goworkflowdocker/utils"

	"github.com/docker/docker/client"
)

type UploadController struct {
	url         string
	cli         *client.Client
	identity    string
	timeout     string
	path        string
	addresses   string
	connections []string
}

func NewUploadController(cli *client.Client, identity, timeout string, connections []string) *UploadController {
	fmt.Println("tata1")
	uploadController := new(UploadController)
	fmt.Println("tata2")
	uploadController.cli = cli
	fmt.Println("tata3")

	uploadController.identity = identity
	fmt.Println("tata4")

	uploadController.timeout = timeout
	fmt.Println("tata5")

	uploadController.connections = connections
	fmt.Println("tata6")

	uploadController.path = utils.GetPathFromConnections(connections)
	fmt.Println("tata7")

	uploadController.addresses = utils.ChangePathFromConnections(connections)
	fmt.Println("tata8")

	return uploadController
}

func (uc UploadController) Get(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET")
	tmpl := template.New("name")
	tmpl = template.Must(tmpl.ParseFiles("upload/tmpl/layout.tmpl", "upload/tmpl/content.tmpl"))

	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "App"})

}

func (uc UploadController) Post(w http.ResponseWriter, r *http.Request) {

	fmt.Println("POST")
	err := r.ParseForm()

	if err != nil {
		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("upload/tmpl/layout.tmpl", "upload/tmpl/content.tmpl"))
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "App"})
	} else {

		fmt.Println("File Upload Endpoint Hit")

		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		name := r.FormValue("name")
		fmt.Println("name")
		fmt.Println(name)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		path := createDirectory()
		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		// Create file
		dst, err := os.Create(path + "/" + handler.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := copyFiles(path); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := buildImage(uc.cli, name, path); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := runImage(uc.cli, uc.identity, uc.timeout, uc.addresses, name, uc.path); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}

}

func runImage(cli *client.Client, identity, timeout, addresses, name, path string) error {
	inputEnv := []string{fmt.Sprintf("env_identity=%s", identity), fmt.Sprintf("env_timeout=%s", timeout), fmt.Sprintf("env_adresses=%s", addresses)}
	_, err := docker.RunContainer(cli, name, path, inputEnv)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func buildImage(cli *client.Client, name, path string) error {
	uid := strconv.Itoa(os.Getuid())
	gid := strconv.Itoa(os.Getgid())

	args := make(map[string]*string)
	args["gandalf_uid"] = &uid
	args["gandalf_gid"] = &gid

	tags := []string{name}
	err := docker.BuildImage(cli, tags, path, args)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func createDirectory() string {
	path := "/tmp/workflow_" + time.Now().Format("2006-01-02 15:04:05")
	err := os.Mkdir(path, 0755)
	if err == nil {
		return path
	} else {
		return ""
	}
}

func copyFiles(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	src := dir + "/files"
	dst := path

	if err = copy(src+"/main.go", dst+"/main.go"); err != nil {
		return err
	}

	if err = copy(src+"/Dockerfile", dst+"/Dockerfile"); err != nil {
		return err
	}

	if err = copy(src+"/go.mod", dst+"/go.mod"); err != nil {
		return err
	}

	if err = copy(src+"/setup.sh", dst+"/setup.sh"); err != nil {
		return err
	}

	return nil
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
