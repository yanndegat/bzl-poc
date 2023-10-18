package main

import (
	"log"
	"os"

	"components/go-hello/cmd"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	if err := cmd.Execute(); err != nil {
		log.Printf("[ERROR] Failed: %v", err)
		return -1
	}

	return 0
}
