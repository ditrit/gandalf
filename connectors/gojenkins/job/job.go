package job

import (
	"context"
	"fmt"
	"time"

	"github.com/bndr/gojenkins"
)

type GetBuildPayload struct {
	JobName string
	Number  int64
}

func GetBuild(client *gojenkins.Jenkins, jobName string, number int64) (build *gojenkins.Build, err error) {
	ctx := context.Background()
	build, err = client.GetBuild(ctx, jobName, number)

	return
}

type BuildJobPayload struct {
	Username string
	Password string
	URL      string
	JobName  string
	Params   map[string]string
}

func BuildJob(client *gojenkins.Jenkins, name string, params map[string]string) (int64, error) {
	ctx := context.Background()
	return client.BuildJob(ctx, name, params)
}

func GetJobObj(client *gojenkins.Jenkins, jobName string, number int64) (build *gojenkins.Build, err error) {
	ctx := context.Background()
	build, err = client.GetBuild(ctx, jobName, number)

	return
}

func GetAllBuildIds(client *gojenkins.Jenkins, jobName string) (build *gojenkins.Build, err error) {
	ctx := context.Background()
	build, err = client.GetAllBuildIds(ctx, jobName)

	return
}

type GetBuildFromQueueIDPayload struct {
	JobName string
	Queueid int64
}

func GetBuildFromQueueID(client *gojenkins.Jenkins, jobName string, queueid int64) (int64, string, error) {
	ctx := context.Background()

	job := client.GetJobObj(ctx, jobName)

	build, err := client.GetBuildFromQueueID(ctx, job, queueid)
	if err != nil {
		panic(err)
	}

	// Wait for build to finish
	for build.IsRunning(ctx) {
		time.Sleep(5000 * time.Millisecond)
		build.Poll(ctx)
	}

	fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())

	return build.GetBuildNumber(), build.GetResult(), err
}

type GetLastSuccessfulBuildPayload struct {
	Username string
	Password string
	URL      string
	JobName  string
}

func GetLastSuccessfulBuild(client *gojenkins.Jenkins, jobName string) (int64, string, error) {
	ctx := context.Background()

	job := client.GetJobObj(ctx, jobName)

	build, err := job.GetLastSuccessfulBuild(ctx)
	if err == nil {
		return build.GetBuildNumber(), build.GetResult(), err
	}

	return -1, "", err
}

type GetLastStableBuildPayload struct {
	JobName string
}

func GetLastStableBuild(client *gojenkins.Jenkins, jobName string) (int64, string, error) {
	ctx := context.Background()

	job := client.GetJobObj(ctx, jobName)

	build, err := job.GetLastStableBuild(ctx)

	return build.GetBuildNumber(), build.GetResult(), err

}
