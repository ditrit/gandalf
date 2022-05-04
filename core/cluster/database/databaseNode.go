package database

import (
	"fmt"
	"os"
	"os/exec"
)

func CoackroachStart(dataDir, certsDir, node, bindAddress, httpAddress, members string) error {
	fmt.Println(dataDir)
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
	cmd := exec.Command("/usr/local/bin/cockroach", "init", "--certs-dir="+certsDir, "--host=127.0.0.1:9299")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	fmt.Println("stop2")

	return err
}

func CoackroachCreateDatabase(certsDir, host, database, password string) error {
	cmd := exec.Command("/usr/local/bin/cockroach", "sql", "--certs-dir="+certsDir, "--host=127.0.0.1:9299", "--execute=CREATE DATABASE IF NOT EXISTS "+database+"; CREATE USER IF NOT EXISTS "+database+" WITH PASSWORD '"+password+"'; GRANT ALL ON DATABASE "+database+" TO "+database+";")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	err := cmd.Wait()

	fmt.Println("stop3")
	return err
}
