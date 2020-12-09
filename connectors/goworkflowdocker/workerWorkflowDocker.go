package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"shoset/msg"

	"github.com/ditrit/gandalf/connectors/goworkflowdocker/payload"
	"github.com/ditrit/gandalf/libraries/goclient"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//main : main
func main() {

	var major = int64(1)
	var minor = int64(5)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)

	worker.RegisterCommandsFuncs("REGISTER", Register)
	worker.RegisterCommandsFuncs("EXECUTE", Execute)

	worker.Run()
}

func Register(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var registerPayload payload.RegisterPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &registerPayload)
	if err == nil {
		register(registerPayload.GetName(), registerPayload.GetContent())
		return 0
	}

	return 1
}

func Execute(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var executePayload payload.ExecutePayload
	err := json.Unmarshal([]byte(command.GetPayload()), &executePayload)
	if err == nil {
		execute(executePayload.GetName())
		return 0
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

func execute(name string) {
	go func() {
		cmd := exec.Command("docker", "run", "-d", "-it", "--name", name, "--mount", "type=bind,source=/var/run/gandalf,target=/var/Run/gandalf")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Docker exec failed with %s\n", err)
		}
	}()
}
