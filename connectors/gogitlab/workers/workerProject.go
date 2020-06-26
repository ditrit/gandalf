package workers

import (
	"connectors/gogitlab/client"
	"connectors/gogitlab/project"
	"encoding/json"
	"fmt"
	"libraries/goclient"
	"libraries/goclient/models"
)

type WorkerProject struct {
	ClientGandalf *goclient.ClientGandalf
	ClientGilab   *client.ClientGitlab
	WorkerServer  *WorkerServer
}

func NewWorkerProject(identity, token string, connections []string, workerServer *WorkerServer) *WorkerProject {
	workerProject := new(WorkerProject)
	workerProject.ClientGandalf = goclient.NewClientGandalf(identity, "", connections)
	workerProject.ClientGilab = client.NewClient(token)
	workerProject.WorkerServer = workerServer

	return workerProject
}

func NewWorkerProjectBasic(identity, login, password, rooturl string, connections []string, workerServer *WorkerServer) *WorkerProject {
	workerProject := new(WorkerProject)
	workerProject.ClientGandalf = goclient.NewClientGandalf(identity, "", connections)
	workerProject.ClientGilab = client.NewClientBasic(login, password, rooturl)
	workerProject.WorkerServer = workerServer

	return workerProject
}

func NewWorkerProjectTEST(login, password, rooturl string, workerServer *WorkerServer) *WorkerProject {
	workerProject := new(WorkerProject)
	fmt.Println(client.NewClientBasic(login, password, rooturl))
	workerProject.ClientGilab = client.NewClientBasic(login, password, rooturl)
	workerProject.WorkerServer = workerServer

	return workerProject
}

func (r WorkerProject) Run() {
	done := make(chan bool)
	go r.AddProjectHook()
	go r.ListProjectHook()
	go r.ListProject()
	<-done
}

func (r WorkerProject) AddProjectHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("ADD_PROJECT_HOOK", id)
		fmt.Println(command)

		var addProjectHookPayload project.AddProjectHookPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &addProjectHookPayload)

		if err == nil {
			result := project.AddProjectHook(r.ClientGilab, *addProjectHookPayload.Pid, r.WorkerServer.GetUrl(), *addProjectHookPayload.Token, addProjectHookPayload.ConfidentialNoteEvents,
				addProjectHookPayload.PushEvents, addProjectHookPayload.IssuesEvents, addProjectHookPayload.ConfidentialIssuesEvents,
				addProjectHookPayload.MergeRequestsEvents, addProjectHookPayload.TagPushEvents, addProjectHookPayload.NoteEvents,
				addProjectHookPayload.JobEvents, addProjectHookPayload.PipelineEvents, addProjectHookPayload.WikiPageEvents,
				addProjectHookPayload.EnableSSLVerification)
			if result {
				r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
			} else {
				r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
			}
		}
	}
}

func (r WorkerProject) AddProjectHookServer() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("ADD_PROJECT_HOOK_SERVER", id)
		fmt.Println(command)

		var addProjectHookPayload project.AddProjectHookServerPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &addProjectHookPayload)

		if err == nil {
			result := project.AddProjectHook(r.ClientGilab, *addProjectHookPayload.Pid, *addProjectHookPayload.URL, *addProjectHookPayload.Token, addProjectHookPayload.ConfidentialNoteEvents,
				addProjectHookPayload.PushEvents, addProjectHookPayload.IssuesEvents, addProjectHookPayload.ConfidentialIssuesEvents,
				addProjectHookPayload.MergeRequestsEvents, addProjectHookPayload.TagPushEvents, addProjectHookPayload.NoteEvents,
				addProjectHookPayload.JobEvents, addProjectHookPayload.PipelineEvents, addProjectHookPayload.WikiPageEvents,
				addProjectHookPayload.EnableSSLVerification)
			if result {
				r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
			} else {
				r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
			}
		}
	}
}

func (r WorkerProject) AddProjectHookTEST(pid, token string, mergeRequestsEvents,
	confidentialNoteEvents, pushEvents, issuesEvents, confidentialIssuesEvents, tagPushEvents, noteEvents, jobEvents,
	pipelineEvent, wikiPageEvents, enableSSLVerification bool) {

	result := project.AddProjectHook(r.ClientGilab, pid, r.WorkerServer.GetUrl(), token, &mergeRequestsEvents, &confidentialNoteEvents, &pushEvents,
		&issuesEvents, &confidentialIssuesEvents, &tagPushEvents, &noteEvents, &jobEvents, &pipelineEvent, &wikiPageEvents, &enableSSLVerification)

	fmt.Println("result")
	fmt.Println(result)

}

func (r WorkerProject) ListProjectHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("LIST_PROJECT_HOOK", id)
		fmt.Println(command)

		var listProjectHooksPayload project.ListProjectHooksPayload
		json.Unmarshal([]byte(command.GetPayload()), &listProjectHooksPayload)

		projectHooks, result := project.ListProjectHooks(r.ClientGilab, *listProjectHooksPayload.Pid)
		fmt.Println(projectHooks)
		if result {
			r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
		} else {
			r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
		}
	}
}

func (r WorkerProject) ListProject() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("LIST_PROJECT", id)
		fmt.Println(command)

		projects, result := project.ListProjects(r.ClientGilab)
		fmt.Println(projects)
		if result {
			r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
		} else {
			r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
		}
	}
}
