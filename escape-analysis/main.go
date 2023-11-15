package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("go package path required\n")
		return
	}

	path := os.Args[1]

	var stderr bytes.Buffer
	cmd := exec.Command("go", "build", "-gcflags=-m", path)
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("failed to execute command %s\n", err)
		return
	}

	analysis, err := NewEscapeAnalysis()
	if err != nil {
		fmt.Printf("failed to instantiate a new escape analysis %s\n", err)
		return
	}

	analysis.Run(stderr.String())
	fmt.Println(analysis)
}
