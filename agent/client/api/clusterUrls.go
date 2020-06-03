package api

type urls struct {
	STATIC_PATH string

	GANDALF_PATH string

	CLUSTER_PATH        string
	CLUSTER_PATH_INDEX  string
	CLUSTER_PATH_CREATE string
	CLUSTER_PATH_UPDATE string
	CLUSTER_PATH_DELETE string
}

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.LOGIN_PATH = "/"
	url_patterns.LOGOUT_PATH = "/logout"
	url_patterns.GANDALF_PATH = "/admin/"

	//CLUSTER
	url_patterns.CLUSTER_PATH = url_patterns.GANDALF_PATH + "cluster/"
	url_patterns.CLUSTER_PATH_INDEX = url_patterns.CLUSTER_PATH + "index/"
	url_patterns.CLUSTER_PATH_CREATE = url_patterns.CLUSTER_PATH + "create/"
	url_patterns.CLUSTER_PATH_UPDATE = url_patterns.CLUSTER_PATH + "update/:id/"
	url_patterns.CLUSTER_PATH_DELETE = url_patterns.CLUSTER_PATH + "delete/:id/"

	return url_patterns
}
