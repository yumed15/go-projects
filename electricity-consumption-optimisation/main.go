package main

import "fmt"

func main() {
	err := process()
	if err != nil {
		fmt.Print(err)
	}
}
