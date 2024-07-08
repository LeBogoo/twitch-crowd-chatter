package modules

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
)

func init() {
	fmt.Println("Collect module loaded")
	registerCommand("collect", Collect)
}

func Collect() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: twitch-crowd-chatter collect <channelName>")
		return
	}

	var channelName = os.Args[2]
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// append message to a file called messages.txt
		f, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		// append time, username and message content to the file and the console
		fmt.Println(message.Time.Format("2006-01-02 15:04:05"), message.User.Name, message.Message)
		if _, err := f.WriteString(message.Time.Format("2006-01-02 15:04:05") + " " + message.User.Name + ": " + message.Message + "\n"); err != nil {
			fmt.Println(err)
			return
		}

	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
