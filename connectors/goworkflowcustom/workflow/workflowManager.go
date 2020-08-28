package workflow

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func ExecuteWorkflow(filepath, filename string) {
	fileNameWithtoutExtension := strings.TrimSuffix(filename, path.Ext(filename))
	buildWorkflow(filepath, fileNameWithtoutExtension)
	runWorkflow(filepath, fileNameWithtoutExtension)
}

func buildWorkflow(filepath, filename string) {
	fmt.Println("Build workflow")

	cmd := exec.Command("go", "build", "-o", filename)
	cmd.Dir = filepath
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Run() failed with %s\n", err)
	} else {
		fmt.Println("cmd.Run() done")
		fmt.Printf("Output: %q\n", out.String())
	}
}

func runWorkflow(filepath, filename string) {
	fmt.Println("Starting workflow")

	cmd := exec.Command("./" + filename)
	cmd.Dir = filepath
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Run() failed with %s\n", err)
	} else {
		fmt.Println("cmd.Run() done")
		fmt.Printf("Output: %q\n", out.String())
	}

}
