package app

type urls struct {
	STATIC_PATH   string
	GANDALF_PATH  string
	APP_PATH      string
	APP_PATH_GET  string
	APP_PATH_POST string
}

func ReturnURLS() string {
	return "/gandalf/app"
}
