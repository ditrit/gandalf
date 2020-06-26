package project

import (
	"connectors/gogitlab/client"
	"log"

	"github.com/xanzy/go-gitlab"
)

func ListProjects(clientGitlab *client.ClientGitlab) ([]*gitlab.Project, bool) {
	var result = true
	projects, _, err := clientGitlab.Client.Projects.ListProjects(nil)
	if err != nil {
		result = false
	}

	log.Printf("Found %d projects", len(projects))

	return projects, result
}

func ListProjectHooks(clientGitlab *client.ClientGitlab, pid string) ([]*gitlab.ProjectHook, bool) {
	var result = true
	projects, _, err := clientGitlab.Client.Projects.ListProjectHooks(pid, nil, nil)
	if err != nil {
		result = false
	}

	log.Printf("Found %d projects", len(projects))

	return projects, result
}

func AddProjectHook(clientGitlab *client.ClientGitlab, pid, url, token string, mergeRequestsEvents, confidentialNoteEvents, pushEvents, issuesEvents, confidentialIssuesEvents, tagPushEvents, noteEvents, jobEvents,
	pipelineEvent, wikiPageEvents, enableSSLVerification *bool) bool {

	result := true
	opt := &gitlab.AddProjectHookOptions{URL: &url, ConfidentialNoteEvents: confidentialNoteEvents,
		PushEvents: pushEvents, IssuesEvents: issuesEvents, ConfidentialIssuesEvents: confidentialIssuesEvents,
		MergeRequestsEvents: mergeRequestsEvents, TagPushEvents: tagPushEvents, NoteEvents: noteEvents, JobEvents: jobEvents,
		PipelineEvents: pipelineEvent, WikiPageEvents: wikiPageEvents, EnableSSLVerification: enableSSLVerification, Token: &token}

	_, _, err := clientGitlab.Client.Projects.AddProjectHook(pid, opt)

	if err != nil {
		result = false
	}

	return result
}
