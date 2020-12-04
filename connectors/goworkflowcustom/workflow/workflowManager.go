package workflow

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func ExecuteWorkflow(filepath, filename string) {
	fileNameWithtoutExtension := strings.TrimSuffix(filename, path.Ext(filename))
	buildWorkflow(filepath, fileNameWithtoutExtension)
	//executableWorkflow(filepath, fileNameWithtoutExtension)

	runWorkflow(filepath, fileNameWithtoutExtension)
}

func buildWorkflow(filepath, filename string) {
	fmt.Println("Build workflow")

	cmd := exec.Command("go", "build", "-o", filename)
	cmd.Path = filepath
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
	/* cmd.Dir = filepath
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Run() failed with %s\n", err)
	} else {
		fmt.Println("cmd.Run() done")
		fmt.Printf("Output: %q\n", out.String())
	} */
}

func runWorkflow(filepath, filename string) {
	fmt.Println("Starting workflow")
	fmt.Println("filename")
	fmt.Println(filename)
	fmt.Println("filepath")
	fmt.Println(filepath)

	cmd := exec.Command("./" + filename)
	cmd.Path = filepath
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
	/* var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Run() failed with %s\n", err)
	} else {
		fmt.Println("cmd.Run() done")
		fmt.Printf("Output: %q\n", out.String())
	} */

}
