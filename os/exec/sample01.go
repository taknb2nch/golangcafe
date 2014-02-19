package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
)

func main() {
	path, err := exec.LookPath("cmd")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	fmt.Printf("Path = %s\n", path)

	cmd := exec.Command("cmd", "/c", "dir", "c:\\golang\\go\\")

	stdoutpipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("StdoutPipe Error: %v\n", err)
		return
	}
	defer stdoutpipe.Close()

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Command Start Error: %v\n", err)
		return
	}

	fmt.Println("コマンドを起動しました。")

	stdout, err := ioutil.ReadAll(
		transform.NewReader(stdoutpipe, japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		fmt.Printf("Command Error: %v\n", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command Wait Error: %v\n", err)
		return
	}

	fmt.Printf("%s\n", stdout)
}
