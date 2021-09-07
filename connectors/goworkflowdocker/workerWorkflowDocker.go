package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/connectors/goworkflowdocker/docker"
	"github.com/ditrit/gandalf/connectors/goworkflowdocker/upload"
	"github.com/ditrit/gandalf/connectors/goworkflowdocker/utils"

	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/ditrit/shoset/msg"
	"github.com/docker/docker/client"

	payload "github.com/ditrit/gandalf/connectors/goworkflowdocker/payload"
	"github.com/ditrit/gandalf/libraries/goclient"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//main : main
func main() {

	var major = int64(1)
	var minor = int64(0)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("ERR")
		log.Fatalln("Unable to create docker client")
	}

	fmt.Println("identity")
	fmt.Println(worker.GetIdentity())
	fmt.Println("timeout")
	fmt.Println(worker.GetTimeout())
	fmt.Println("connections")
	fmt.Println(worker.GetConnections())

	worker.Context["client"] = cli
	worker.Context["identity"] = worker.GetIdentity()
	worker.Context["timeout"] = worker.GetTimeout()
	worker.Context["connections"] = worker.GetConnections()

	//worker.RegisterCommandsFuncs("REGISTER", Register)
	//worker.RegisterCommandsFuncs("EXECUTE", Execute)
	fmt.Println("START")
	serverupload := upload.NewServerUpload(cli, worker.GetIdentity(), worker.GetTimeout(), worker.GetConnections())
	fmt.Println("START2")
	worker.RegisterServicesFuncs("UploadService", serverupload.Run)

	fmt.Println("RUN")
	worker.Run()
}

func Register(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var registerPayload payload.RegisterPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &registerPayload)
	if err == nil {
		register(registerPayload.GetName(), registerPayload.GetContent())
		return 0
	}

	return 1
}

func Execute(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	cli := context["client"].(*client.Client)
	identity := context["identity"].(string)
	timeout := context["timeout"].(string)
	connections := context["connections"].([]string)
	path := utils.GetPathFromConnections(connections)
	addresses := utils.TransformConnectionsSliceToString(connections)
	var executePayload payload.ExecutePayload
	err := json.Unmarshal([]byte(command.GetPayload()), &executePayload)
	if err == nil {
		err = execute(cli, identity, timeout, addresses, executePayload.GetName(), path)
		if err == nil {
			return 0
		}
	}

	return 1
}

func register(name, content string) {
	// Récupérer le user courant
	user, err := user.Current()
	if err != nil {
		log.Fatal("Unable to find current user")
	}

	// Récupérer le répertoire courant
	wDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to find current working directory")
	}

	// Verifier que les arguments existent
	if name == "" {
		log.Fatal("No name provided")
	}
	if content == "" {
		log.Fatal("No golang workflow provided")
	}

	// Repertoire de build
	dstDir := filepath.Join(user.HomeDir, name)

	// Suppresion des anciennes versions du repertoire de build
	_, err = os.Stat(dstDir)
	if err == nil {
		err := os.RemoveAll(dstDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Creation du repertoire de build
	err = os.MkdirAll(dstDir, 0777)
	if err != nil {
		log.Fatal("Unable to create destination directory")
	}

	// Mise en place du Dockerfile
	err = os.Link(
		filepath.Join(wDir, "Dockerfile"),
		filepath.Join(dstDir, "Dockerfile"))
	if err != nil {
		log.Fatal("Unable to create Dockerfile")
	}

	// Ecriture du fichier de workflow
	wkfPath := filepath.Join(dstDir, name+".go")
	wkfFile, err := os.Create(wkfPath)
	if err != nil {
		log.Fatal("Unable to write workflow file")
	}
	defer wkfFile.Close()
	_, err = wkfFile.WriteString(content)

	// Deplacement dans le repertoire de travail
	err = os.Chdir(dstDir)
	if err != nil {
		log.Fatal("can not change directory to" + dstDir)
	}

	// Construction du conteneur
	cmd := exec.Command("docker", "build", "-t", name, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Docker build failed with %s\n", err)
	}

	// Replacement dans le repertoire de travail initial
	os.Chdir(wDir)

}

func execute(cli *client.Client, identity, timeout, connections, name, path string) error {
	inputEnv := []string{fmt.Sprintf("env_identity=%s", identity), fmt.Sprintf("env_timeout=%s", timeout), fmt.Sprintf("env_adresses=%s", connections)}
	_, err := docker.RunContainer(cli, name, path, inputEnv)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
