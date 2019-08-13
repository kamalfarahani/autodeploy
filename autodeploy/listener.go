package autodeploy

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// RegisterListeners makes endpoints for given listeners
func RegisterListeners(endPointToScriptPath map[string]string) {
	for endPoint, scriptPath := range endPointToScriptPath {
		err := validatePath(scriptPath)
		if err != nil {
			panic(err)
		}

		http.HandleFunc(endPoint, makeListener(scriptPath))
	}
}

func makeListener(scriptPath string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		err := executeFile(scriptPath)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		} else {
			fmt.Fprintf(writer, "OK")
		}
	}
}

func executeFile(scriptPath string) error {
	cmd := exec.Command("/bin/sh", scriptPath)

	return cmd.Run()
}

func validatePath(path string) error {
	_, err := os.Stat(path)

	return err
}
