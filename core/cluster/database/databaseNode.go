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

func CoackroachStart(dataDir, certsDir, node, bindAddress, httpAddress, members string) error {
	fmt.Println(dataDir)
	//path, err := os.Getwd()
	//cmdString := "cockroach start --certs-dir=" + certsDir + " --store=" + node + " --listen-addr=" + bindAddress + " --http-addr=" + httpAddress + " --join=" + members + " --background"
	//cmd := exec.Command(cmdString)
	cmd := exec.Command("/usr/local/bin/cockroach", "start", "--certs-dir="+certsDir, "--store="+node, "--listen-addr="+bindAddress, "--http-addr="+httpAddress, "--join="+members, "--background")
	cmd.Dir = dataDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	fmt.Println("stop")
	return err
}

func CoackroachInit(certsDir, host string) error {
	//path, err := os.Getwd()
	//cmdString := "/usr/local/bin/cockroach init --certs-dir=" + certsDir + " --host=" + host
	cmd := exec.Command("/usr/local/bin/cockroach", "init", "--certs-dir="+certsDir, "--host="+host)
	//cmd.Dir = path + "/cluster/database/"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	fmt.Println("stop2")

	return err
}

func CoackroachCreateDatabase(certsDir, host, database, password string) error {

	//path, err := os.Getwd()
	//cmdString := "/usr/local/bin/cockroach sql 	--certs-dir=" + certsDir + " --host=" + host + " <<EOF CREATE DATABASE IF NOT EXISTS " + database + "; CREATE USER IF NOT EXISTS " + database + " WITH PASSWORD '" + password + "'; GRANT ALL ON DATABASE " + database + " TO " + database + "; EOF"
	cmd := exec.Command("/usr/local/bin/cockroach", "sql", "--certs-dir="+certsDir, "--host="+host, "--execute=CREATE DATABASE IF NOT EXISTS "+database+"; CREATE USER IF NOT EXISTS "+database+" WITH PASSWORD '"+password+"'; GRANT ALL ON DATABASE "+database+" TO "+database+";")
	//cmd.Dir = path + "/cluster/database/"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	err := cmd.Wait()

	fmt.Println("stop3")
	return err
}
