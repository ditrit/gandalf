package api

type urls struct {
	STATIC_PATH            string
	GANDALF_PATH           string
	AGGREGATOR_PATH        string
	AGGREGATOR_PATH_LIST   string
	AGGREGATOR_PATH_CREATE string
	AGGREGATOR_PATH_READ   string
	AGGREGATOR_PATH_UPDATE string
	AGGREGATOR_PATH_DELETE string
	CLUSTER_PATH           string
	CLUSTER_PATH_LIST      string
	CLUSTER_PATH_CREATE    string
	CLUSTER_PATH_READ      string
	CLUSTER_PATH_UPDATE    string
	CLUSTER_PATH_DELETE    string
	CONNECTOR_PATH         string
	CONNECTOR_PATH_LIST    string
	CONNECTOR_PATH_CREATE  string
	CONNECTOR_PATH_READ    string
	CONNECTOR_PATH_UPDATE  string
	CONNECTOR_PATH_DELETE  string
	ROLE_PATH              string
	ROLE_PATH_LIST         string
	ROLE_PATH_CREATE       string
	ROLE_PATH_READ         string
	ROLE_PATH_UPDATE       string
	ROLE_PATH_DELETE       string
	USER_PATH              string
	USER_PATH_LIST         string
	USER_PATH_CREATE       string
	USER_PATH_READ         string
	USER_PATH_UPDATE       string
	USER_PATH_DELETE       string
	TENANT_PATH            string
	TENANT_PATH_LIST       string
	TENANT_PATH_CREATE     string
	TENANT_PATH_READ       string
	TENANT_PATH_UPDATE     string
	TENANT_PATH_DELETE     string
}

func ReturnURLS() *urls {

	apiurls := new(urls)
	apiurls.GANDALF_PATH = "/"
	apiurls.AGGREGATOR_PATH = "/aggregator"
	apiurls.AGGREGATOR_PATH_LIST = apiurls.AGGREGATOR_PATH + "/list"
	apiurls.AGGREGATOR_PATH_CREATE = apiurls.AGGREGATOR_PATH + "/create"
	apiurls.AGGREGATOR_PATH_READ = apiurls.AGGREGATOR_PATH + "/read/{id}"
	apiurls.AGGREGATOR_PATH_UPDATE = apiurls.AGGREGATOR_PATH + "/update/{id}"
	apiurls.AGGREGATOR_PATH_DELETE = apiurls.AGGREGATOR_PATH + "/delete/{id}"
	apiurls.CLUSTER_PATH = "/cluster"
	apiurls.CLUSTER_PATH_LIST = apiurls.CLUSTER_PATH + "/list"
	apiurls.CLUSTER_PATH_CREATE = apiurls.CLUSTER_PATH + "/create"
	apiurls.CLUSTER_PATH_READ = apiurls.CLUSTER_PATH + "/read/{id}"
	apiurls.CLUSTER_PATH_UPDATE = apiurls.CLUSTER_PATH + "/update/{id}"
	apiurls.CLUSTER_PATH_DELETE = apiurls.CLUSTER_PATH + "/delete/{id}"
	apiurls.CONNECTOR_PATH = "/connector"
	apiurls.CONNECTOR_PATH_LIST = apiurls.CONNECTOR_PATH + "/list"
	apiurls.CONNECTOR_PATH_CREATE = apiurls.CONNECTOR_PATH + "/create"
	apiurls.CONNECTOR_PATH_READ = apiurls.CONNECTOR_PATH + "/read/{id}"
	apiurls.CONNECTOR_PATH_UPDATE = apiurls.CONNECTOR_PATH + "/update/{id}"
	apiurls.CONNECTOR_PATH_DELETE = apiurls.CONNECTOR_PATH + "/delete/{id}"
	apiurls.ROLE_PATH = "/role"
	apiurls.ROLE_PATH_LIST = apiurls.ROLE_PATH + "/list"
	apiurls.ROLE_PATH_CREATE = apiurls.ROLE_PATH + "/create"
	apiurls.ROLE_PATH_READ = apiurls.ROLE_PATH + "/read/{id}"
	apiurls.ROLE_PATH_UPDATE = apiurls.ROLE_PATH + "/update/{id}"
	apiurls.ROLE_PATH_DELETE = apiurls.ROLE_PATH + "/delete/{id}"
	apiurls.USER_PATH = "/user"
	apiurls.USER_PATH_LIST = apiurls.USER_PATH + "/list"
	apiurls.USER_PATH_CREATE = apiurls.USER_PATH + "/create"
	apiurls.USER_PATH_READ = apiurls.USER_PATH + "/read/{id}"
	apiurls.USER_PATH_UPDATE = apiurls.USER_PATH + "/update/{id}"
	apiurls.USER_PATH_DELETE = apiurls.USER_PATH + "/delete/{id}"
	apiurls.TENANT_PATH = "/tenant"
	apiurls.TENANT_PATH_LIST = apiurls.TENANT_PATH + "/list"
	apiurls.TENANT_PATH_CREATE = apiurls.TENANT_PATH + "/create"
	apiurls.TENANT_PATH_READ = apiurls.TENANT_PATH + "/read/{id}"
	apiurls.TENANT_PATH_UPDATE = apiurls.TENANT_PATH + "/update/{id}"
	apiurls.TENANT_PATH_DELETE = apiurls.TENANT_PATH + "/delete/{id}"

	return apiurls
}
