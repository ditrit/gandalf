package database

import (
	"fmt"
	"os"
	"os/exec"
)

/* func installCockroach() {

	cmd := exec.Command("/bin/sh", "./installCockroach.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Printf("Can't install cockroach")
	}
}

func setupCoackroach() {
	cmd := exec.Command("/bin/sh", "./setupCockroach.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Printf("Can't setup cockroach")
	}
} */

func CoackroachStart(dataDir, node, bindAddress, httpAddress, members string) error {
	path, err := os.Getwd()
	fmt.Println(path)
	cmd := exec.Command("/bin/sh", "./cockroachStart.sh", dataDir, node, bindAddress, httpAddress, members)
	cmd.Dir = path + "/cluster/database/"
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Start()
	err = cmd.Wait()

	return err
}

func CoackroachInit(dataDir, host string) error {
	path, err := os.Getwd()
	fmt.Println(path)
	cmd := exec.Command("/bin/sh", "./cockroachInit.sh", dataDir, host)
	cmd.Dir = path + "/cluster/database/"

	cmd.Start()
	err = cmd.Wait()
	return err
}

func CoackroachCreateDatabase(dataDir, host, database string) error {
	path, err := os.Getwd()
	fmt.Println(path)
	cmd := exec.Command("/bin/sh", "./cockroachCreateDatabase.sh", dataDir, host, database)
	cmd.Dir = path + "/cluster/database/"

	cmd.Start()
	err = cmd.Wait()
	return err
}
