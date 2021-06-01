package database

import (
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

func CoackroachStart(dataDir, certsDir, node, bindAddress, httpAddress, members string) error {
	path, err := os.Getwd()
	cmd := exec.Command("/bin/sh", "./cockroachStart.sh", dataDir, certsDir, node, bindAddress, httpAddress, members)
	cmd.Dir = path + "/cluster/database/"

	cmd.Start()
	err = cmd.Wait()
	return err
}

func CoackroachInit(certsDir, host string) error {
	path, err := os.Getwd()
	cmd := exec.Command("/bin/sh", "./cockroachInit.sh", certsDir, host)
	cmd.Dir = path + "/cluster/database/"

	cmd.Start()
	err = cmd.Wait()
	return err
}

func CoackroachCreateDatabase(certsDir, host, database, password string) error {
	path, err := os.Getwd()
	cmd := exec.Command("/bin/sh", "./cockroachCreateDatabase.sh", certsDir, host, database, password)
	cmd.Dir = path + "/cluster/database/"

	cmd.Start()
	err = cmd.Wait()
	return err
}
