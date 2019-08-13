package autodeploy

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// Serve start listening for registerd listeners
func Serve(addr string) {
	fmt.Printf("Serving and Listening on %s", addr)
	http.ListenAndServe(addr, nil)
}

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
		result, err := executeFile(scriptPath)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		} else {
			fmt.Fprintf(writer, result)
		}
	}
}

func executeFile(scriptPath string) (string, error) {
	cmd := exec.Command("/bin/sh", scriptPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}

func validatePath(path string) error {
	_, err := os.Stat(path)

	return err
}
