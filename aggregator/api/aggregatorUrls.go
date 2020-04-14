package api

type urls struct {
	STATIC_PATH string

	GANDALF_PATH string

	AGGREGATOR_PATH        string
	AGGREGATOR_PATH_INDEX  string
	AGGREGATOR_PATH_CREATE string
	AGGREGATOR_PATH_UPDATE string
	AGGREGATOR_PATH_DELETE string
}

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.LOGIN_PATH = "/"
	url_patterns.LOGOUT_PATH = "/logout"
	url_patterns.GANDALF_PATH = "/admin/"

	//AGGREGATOR
	url_patterns.AGGREGATOR_PATH = url_patterns.GANDALF_PATH + "aggregator/"
	url_patterns.AGGREGATOR_PATH_INDEX = url_patterns.AGGREGATOR_PATH + "index/"
	url_patterns.AGGREGATOR_PATH_CREATE = url_patterns.AGGREGATOR_PATH + "create/"
	url_patterns.AGGREGATOR_PATH_UPDATE = url_patterns.AGGREGATOR_PATH + "update/:id/"
	url_patterns.AGGREGATOR_PATH_DELETE = url_patterns.AGGREGATOR_PATH + "delete/:id/"

	return url_patterns
}
