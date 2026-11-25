package main

import (
	"fmt"
	"os"

	"github.com/theguuholi/forge/internal/cli"
)

func main() {
	text, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(text)
}
