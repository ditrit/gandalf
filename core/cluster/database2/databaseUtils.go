package database2

import (
	"log"
	"os"
	"os/exec"
)

func installCockroach() {

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
}

func startCoackroach() {
	cmd := exec.Command("/bin/sh", "./startCockroach.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Printf("Can't start cockroach")
	}
}

func initCoackroach() {
	cmd := exec.Command("/bin/sh", "./initCockroach.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Printf("Can't init cockroach")
	}
}
