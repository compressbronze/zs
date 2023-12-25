package main

import (
	"github.com/charmbracelet/log"
	"github.com/satrap-illustrations/zs/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error(err)
	}
}
