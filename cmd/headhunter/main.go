package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/xnet-admin-1/Headhunter/cmd"
)

var version = "dev"

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Printf("headhunter version %s\n", version)
			os.Exit(0)
		}
	}

	rootCmd := cmd.GetRootCommand(version)
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
