package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

type Word struct {
	Text      string
	Frequency int
}

func getPopularMessages(messages []string) map[string]int {
	popularMessages := make(map[string]int)

	for _, message := range messages {
		popularMessages[message]++
	}

	delete(popularMessages, "")

	return popularMessages
}

func getPopularWordsUnfiltered(messages []string) map[string]int {
	popularWords := make(map[string]int)

	for _, message := range messages {
		messageWords := strings.Split(message, " ")
		// make messageWords unique

		for _, word := range messageWords {
			popularWords[word]++
		}
	}

	delete(popularWords, "")

	return popularWords
}

func sortByFrequency(wordsMap map[string]int) []Word {
	var words []Word
	for text, frequency := range wordsMap {
		words = append(words, Word{text, frequency})
	}

	// sort words by frequency
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i].Frequency < words[j].Frequency {
				words[i], words[j] = words[j], words[i]
			}
		}
	}

	return words
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <botName> <botToken> <channelName>")
		return
	}

	var botName = os.Args[1]
	var botToken = os.Args[2]
	var channelName = os.Args[3]
	client := twitch.NewClient(botName, botToken)

	// store the last 50 messages and remove the oldest one if there are more than 50 messages
	var messages []string

	var lastMessageSent = ""
	var lastMessageTime int64 = 0

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		messages = append(messages, message.Message)
		if len(messages) > 50 {
			messages = messages[1:]
		}

		popularWords := getPopularWordsUnfiltered(messages)
		sortedPopularWords := sortByFrequency(popularWords)
		// print the first 3 most popular words if they exist in one line followed by \r
		// clear current line
		fmt.Print("\033[2K")
		fmt.Print("\r")
		fmt.Print("Most popular words: ")
		for i := 0; i < 3 && i < len(sortedPopularWords); i++ {
			fmt.Print(sortedPopularWords[i], " ")
		}

		if sortedPopularWords[0].Frequency > 100 && sortedPopularWords[0].Text == "EDM" {
			// EDM MODE
			// repeat the first two most popular words for 5 times
			var newMessage = ""
			for i := 0; i < 5; i++ {
				newMessage += sortedPopularWords[0].Text + " " + sortedPopularWords[1].Text + " "
			}

			if lastMessageSent != newMessage || time.Now().Unix()-lastMessageTime > 31 {
				fmt.Println()

				fmt.Println("Sending EDM message:", newMessage)
				client.Say(channelName, newMessage)
				lastMessageSent = newMessage
				lastMessageTime = time.Now().Unix()
			}
		} else {
			mostPopularMessages := getPopularMessages(messages)
			sortedPopularMessages := sortByFrequency(mostPopularMessages)

			mostPopularMessage := sortedPopularMessages[0]
			// fmt.Println(mostPopularMessage)
			if mostPopularMessage.Frequency >= 15 && mostPopularMessage.Text != lastMessageSent {
				fmt.Println()
				fmt.Println("Sending message:", mostPopularMessage.Text)
				client.Say(channelName, mostPopularMessage.Text)
				lastMessageSent = mostPopularMessage.Text
			}
		}

	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
