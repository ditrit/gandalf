package form

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

type urls struct {
	STATIC_PATH    string
	GANDALF_PATH   string
	FORM_PATH      string
	FORM_PATH_GET  string
	FORM_PATH_POST string
}

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.GANDALF_PATH = "/gandalf/"

	//FORM
	url_patterns.FORM_PATH_GET = url_patterns.GANDALF_PATH + "form/"
	url_patterns.FORM_PATH_POST = url_patterns.GANDALF_PATH + "form/"
	//url_patterns.FORM_PATH_INDEX = url_patterns.FORM_PATH + "index/"

	return url_patterns
}

func ReturnHashURLS() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	concatenated := fmt.Sprint("/gandalf/form/", random.Intn(100))
	sha_512 := sha512.New()
	sha_512.Write([]byte(concatenated))
	hashurl := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	return "/" + hashurl
}
