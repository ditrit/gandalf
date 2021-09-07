package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func BuildImage(client *client.Client, tags []string, dockerpath string, args map[string]*string) error {
	ctx := context.Background()

	tar, err := archive.TarWithOptions(dockerpath, &archive.TarOptions{})
	if err != nil {
		return err
	}

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Remove:     true,
		Tags:       tags,
		BuildArgs:  args,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		tar,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

func RunContainer(cli *client.Client, imagename, socketpath string, inputEnv []string) (string, error) {
	ctx := context.Background()

	volumes := []string{socketpath + ":/var/run/sockets/:rw"}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image: imagename,
		Env:   inputEnv,
	}
	// Creating the actual container. This is "nil,nil,nil" in every example.
	cont, err := cli.ContainerCreate(
		ctx,
		config,
		&container.HostConfig{
			Binds: volumes,
		},
		nil,
		nil,
		"",
	)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Run the actual container
	if err := cli.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Printf("Container %s is created \n", cont.ID)

	return cont.ID, nil
}

func ListContainers(cli *client.Client) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	fmt.Println("List containers")
	for _, container := range containers {
		fmt.Println(container.ID)
	}
	return nil
}

func StopAllContainers(cli *client.Client) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			return err
		}
		fmt.Println("Success")
	}
	return nil
}

func StopContainer(cli *client.Client, id string) {
	ctx := context.Background()

	if err := cli.ContainerStop(ctx, id, nil); err != nil {
		fmt.Printf("Unable to stop container %s: %s", id, err)
	}
}

func RemoveContainer(cli *client.Client, id string) error {
	ctx := context.Background()

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := cli.ContainerRemove(ctx, id, removeOptions); err != nil {
		fmt.Printf("Unable to remove container: %s", err)
		return err
	}
	return nil
}

func LogContainer(cli *client.Client, id string) error {
	ctx := context.Background()

	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, id, options)
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, out)

	/* 	out, err := cli.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{ShowStdout: true})
	   	if err != nil {
	   		fmt.Println(err)
	   		return err
	   	}

	   	stdcopy.StdCopy(os.Stdout, os.Stderr, out) */

	return nil
}
