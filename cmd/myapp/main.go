package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: myapp <command>")
		fmt.Fprintln(os.Stderr, "commands: serve")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		runServer()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
