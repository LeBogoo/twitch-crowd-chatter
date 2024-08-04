package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	registerCommand("graph", Graph)
}

func Graph() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: twitch-crowd-chatter graph <channelName>")
		return
	}

	var channelName = os.Args[2]

	// read <channelName>.txt line by line because it is really big

	var channelFile = fmt.Sprintf("%s.txt", channelName)
	var file, err = os.Open(channelFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var timeStats = make(map[string]int)

	// read file line by line
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		// 2024-07-04 20:40:55 junal19: plink-182
		line := scanner.Text()

		messageParts := strings.Split(line, " ")
		timeParts := strings.Split(messageParts[1], ":")
		indexString := messageParts[0] + " " + timeParts[0] + ":" + timeParts[1]

		timeStats[indexString]++
	}

	for key, value := range timeStats {
		fmt.Printf("%s, %d\n", key, value)
	}

}
