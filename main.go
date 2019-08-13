package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./autodeploy"
)

var endPointToScript = make(map[string]string)

func main() {
	fillEndPointToScript()
	autodeploy.RegisterListeners(endPointToScript)

	fmt.Print("Enter listen Port: ")
	listenAddr := ":" + readNextLine()
	autodeploy.Serve(listenAddr)
}

func fillEndPointToScript() {
	for {
		fmt.Print("Enter Enpoint: ")
		endPoint := "/" + readNextLine()

		fmt.Print("Enter Script Path: ")
		scriptPath := readNextLine()

		endPointToScript[endPoint] = scriptPath

		fmt.Print("finished?: [y/n]")
		finished := readNextLine()
		if finished == "y" {
			return
		}
	}
}

func readNextLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
