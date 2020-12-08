package server

type urls struct {
	STATIC_PATH       string
	GANDALF_PATH      string
	INDEX_PATH        string
	EVENTS_PATH       string
	SEND_COMMAND_PATH string
	SEND_UPDATE_PATH  string
}

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.GANDALF_PATH = "/gandalf/"

	//FORM
	url_patterns.SEND_COMMAND_PATH = url_patterns.GANDALF_PATH + "cmd/"
	url_patterns.SEND_UPDATE_PATH = url_patterns.GANDALF_PATH + "update/"
	url_patterns.INDEX_PATH = url_patterns.GANDALF_PATH
	url_patterns.EVENTS_PATH = "/events/"
	//url_patterns.FORM_PATH_INDEX = url_patterns.FORM_PATH + "index/"

	return url_patterns
}
