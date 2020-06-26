package project

type ListProjectHooksPayload struct {
	Pid *string
}

type AddProjectHookPayload struct {
	Pid                      *string
	ConfidentialNoteEvents   *bool
	PushEvents               *bool
	IssuesEvents             *bool
	ConfidentialIssuesEvents *bool
	MergeRequestsEvents      *bool
	TagPushEvents            *bool
	NoteEvents               *bool
	JobEvents                *bool
	PipelineEvents           *bool
	WikiPageEvents           *bool
	EnableSSLVerification    *bool
	Token                    *string
}

type AddProjectHookServerPayload struct {
	UrlServer                string
	Pid                      *string
	URL                      *string
	ConfidentialNoteEvents   *bool
	PushEvents               *bool
	IssuesEvents             *bool
	ConfidentialIssuesEvents *bool
	MergeRequestsEvents      *bool
	TagPushEvents            *bool
	NoteEvents               *bool
	JobEvents                *bool
	PipelineEvents           *bool
	WikiPageEvents           *bool
	EnableSSLVerification    *bool
	Token                    *string
}
