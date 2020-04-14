package api

type urls struct {
	STATIC_PATH string

	GANDALF_PATH string

	CONNECTOR_PATH        string
	CONNECTOR_PATH_INDEX  string
	CONNECTOR_PATH_CREATE string
	CONNECTOR_PATH_UPDATE string
	CONNECTOR_PATH_DELETE string
}

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.LOGIN_PATH = "/"
	url_patterns.LOGOUT_PATH = "/logout"
	url_patterns.GANDALF_PATH = "/admin/"

	//CONNECTOR
	url_patterns.CONNECTOR_PATH = url_patterns.GANDALF_PATH + "connector/"
	url_patterns.CONNECTOR_PATH_INDEX = url_patterns.CONNECTOR_PATH + "index/"
	url_patterns.CONNECTOR_PATH_CREATE = url_patterns.CONNECTOR_PATH + "create/"
	url_patterns.CONNECTOR_PATH_UPDATE = url_patterns.CONNECTOR_PATH + "update/:id/"
	url_patterns.CONNECTOR_PATH_DELETE = url_patterns.CONNECTOR_PATH + "delete/:id/"

	return url_patterns
}
