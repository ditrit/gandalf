package upload

import (
	"crypto/sha512"
	"encoding/base64"
)

type urls struct {
	STATIC_PATH  string
	GANDALF_PATH string
	UPLOAD_PATH  string
}

func ReturnURLS() string {
	return "/gandalf/workflow/upload"
}

//TODO REVOIR
func ReturnHashURLS() string {
	sha_512 := sha512.New()
	sha_512.Write([]byte("/gandalf/workflow/upload"))
	hashurl := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	return "/" + hashurl
}
