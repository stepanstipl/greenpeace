package main

import (
	"fmt"
	"io/ioutil"
	"k8s.io/kubernetes/third_party/golang/expansion"
	"os"
	"os/exec"
	"strings"
)

func loadEnv(secretsDir string) {

	files, _ := ioutil.ReadDir(secretsDir)

	for _, f := range files {
		filePath := secretsDir + "/" + f.Name()

		content, _ := ioutil.ReadFile(filePath)
		envValue := strings.TrimSpace(string(content))

		envName := strings.ToUpper(f.Name())

		os.Setenv(envName, envValue)
		// env = append(env, fmt.Sprintf("%s=%s", envName, content_str))
		fmt.Printf("Setting env variable: %s from %s\n", envName, filePath)
	}
}

func expandOsArgs() []string {

	args := []string{}

	if len(os.Args) != 0 {
		for _, cmd := range os.Args {
			args = append(args, expansion.Expand(cmd, os.Getenv))
		}
	}

	return args
}

func runCmd() error {

	args := expandOsArgs()

	cmdCmd := args[1]
	cmdArgs := []string{}
	if len(args) > 2 {
		cmdArgs = args[2:]
	}

	cmd := exec.Command(cmdCmd, cmdArgs...)

	env := os.Environ()
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	return err
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need at least 1 argument - command name")
		os.Exit(1)
	}

	dir := "./secrets"
	loadEnv(dir)

	err := runCmd()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
