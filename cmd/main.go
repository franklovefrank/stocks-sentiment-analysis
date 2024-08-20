package main

import (
	"fmt"
	"os"

	"stock-sentiment-cli/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
