package main

import (
	"fmt"
	"os"
	"strings"
	"twitch-crowd-chatter/modules"
)

func main() {
	var commandNamesList []string
	for name := range modules.Commands {
		commandNamesList = append(commandNamesList, name)
	}

	commandNames := strings.Join(commandNamesList, "|")

	usageMessage := fmt.Sprintf("Usage: twitch-crowd-chatter <%s>", commandNames)

	if len(os.Args) < 2 {
		fmt.Println(usageMessage)
		return
	}

	var subcommand = os.Args[1]
	if command, ok := modules.Commands[subcommand]; ok {
		command()
	} else {
		fmt.Println(usageMessage)
	}

}
