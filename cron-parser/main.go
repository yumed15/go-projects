package main

import (
	"cron-parser/parser"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 7 {
		fmt.Println("cron format needs to have 6 arguments.")
		return
	}

	arguments := os.Args[1:]
	cron := strings.Join(arguments, " ")

	schedule, command, err := parser.Parse(cron)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	output := parser.FormatOutput(schedule, command)
	
	fmt.Println(output)
}
