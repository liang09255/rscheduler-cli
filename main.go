package main

import (
	"rscheduler-cli/cmd"
	"rscheduler-cli/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
