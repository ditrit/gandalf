package upload

type Urls struct {
	ROOT_PATH        string
	PATH             string
	UPLOAD_PATH      string
	UPLOAD_PATH_GET  string
	UPLOAD_PATH_POST string
}

func ReturnURLS() *Urls {
	//BASE
	uploadUrls := new(Urls)
	uploadUrls.ROOT_PATH = "/"
	uploadUrls.PATH = "/gandalf"

	//UPLOAD
	uploadUrls.UPLOAD_PATH = uploadUrls.PATH + "/upload"
	uploadUrls.UPLOAD_PATH_GET = uploadUrls.UPLOAD_PATH + "/"
	uploadUrls.UPLOAD_PATH_POST = uploadUrls.UPLOAD_PATH + "/"

	return uploadUrls
}
