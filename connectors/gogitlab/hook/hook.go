package hook

import (
	"connectors/gogitlab/client"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xanzy/go-gitlab"
)

func ListHook(clientGitlab *client.ClientGitlab) bool {
	result := true
	hooks, err, _ := clientGitlab.Client.SystemHooks.ListHooks()
	if err != nil {
		result = false
	}
	fmt.Println(hooks)
	return result
}

func AddHook(clientGitlab *client.ClientGitlab, url string, pushEvents, tagPushEvents, mergeRequestsEvents, repositoryUpdateEvents, enableSSLVerification *bool) bool {

	result := true
	opt := &gitlab.AddHookOptions{URL: &url, Token: &clientGitlab.Token, PushEvents: pushEvents, TagPushEvents: tagPushEvents, MergeRequestsEvents: mergeRequestsEvents, RepositoryUpdateEvents: repositoryUpdateEvents, EnableSSLVerification: enableSSLVerification}
	hook, err, _ := clientGitlab.Client.SystemHooks.AddHook(opt)

	data := gitlab.Response{}
	json.NewDecoder(err.Body).Decode(&data)
	if err != nil {
		fmt.Println("err")
		fmt.Println(reflect.TypeOf(err))
		fmt.Println(err)

		fmt.Println("data")
		fmt.Println(data)
		result = false
	}
	fmt.Println("hook")
	fmt.Println(hook)
	return result
}

func DeleteHook(clientGitlab *client.ClientGitlab, id int) bool {
	result := true
	err, _ := clientGitlab.Client.SystemHooks.DeleteHook(id, nil)
	if err != nil {
		result = false
	}
	return result
}

func TestHook(clientGitlab *client.ClientGitlab, id int) bool {
	result := true
	hook, err, _ := clientGitlab.Client.SystemHooks.TestHook(id, nil)
	if err != nil {
		result = false
	}
	fmt.Println(hook)

	return result
}
