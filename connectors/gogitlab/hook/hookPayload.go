package hook

type AddHookPayload struct {
	Url                    string
	PushEvents             *bool
	TagPushEvents          *bool
	MergeRequestsEvents    *bool
	RepositoryUpdateEvents *bool
	EnableSSLVerification  *bool
}

type AddHookServerPayload struct {
	UrlServer              string
	UrlHook                string
	PushEvents             *bool
	TagPushEvents          *bool
	MergeRequestsEvents    *bool
	RepositoryUpdateEvents *bool
	EnableSSLVerification  *bool
}

type DeleteHookPayload struct {
	Id int
}

type TestHookPayload struct {
	Id int
}
