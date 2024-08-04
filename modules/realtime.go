package modules

import (
	"fmt"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func init() {
	registerCommand("realtime", Realtime)
}

func Realtime() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: twitch-crowd-chatter realtime <channelName>")
		return
	}

	var channelName = os.Args[2]
	client := twitch.NewAnonymousClient()

	messageCount := 0
	startTime := time.Now()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		f, err := os.OpenFile(channelName+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		// Increment message count
		messageCount++

		// Calculate time since start
		elapsed := time.Since(startTime)

		// Calculate messages per minute
		messagesPerMinute := float64(messageCount) / elapsed.Minutes()

		// Write to console
		fmt.Printf("\rmsgs/min: %.2f", messagesPerMinute)
	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
