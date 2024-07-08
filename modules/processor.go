package modules

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func init() {
	registerCommand("process", Process)
}

func Process() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: twitch-crowd-chatter process <channelName>")
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

	var userStats = make(map[string]int)

	// read file line by line
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		// 2024-07-04 20:40:55 junal19: plink-182
		line := scanner.Text()

		var parts = strings.Split(line, " ")
		var username = strings.Split(parts[2], ":")[0]

		userStats[username]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// sort stats by highest count
	var sortedStats = sortMapByValue(userStats)

	// print top 10
	for i := 0; i < 10; i++ {
		if i >= len(sortedStats) {
			break
		}

		fmt.Printf("#%d %s: %d\n", i+1, sortedStats[i], userStats[sortedStats[i]])
	}

	// print out amount of all chatters
	fmt.Printf("Total chatters: %d\n", len(userStats))
}

func sortMapByValue(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}
