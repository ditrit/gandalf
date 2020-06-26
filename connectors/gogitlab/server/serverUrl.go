package server

import (
	"crypto/sha512"
	"encoding/base64"
)

type urls struct {
	STATIC_PATH  string
	GANDALF_PATH string
	HOOK_PATH    string
}

func ReturnHashURLS() string {
	sha_512 := sha512.New()
	sha_512.Write([]byte("/gandalf/gitlab/"))
	hashurl := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	return "/" + hashurl
}
