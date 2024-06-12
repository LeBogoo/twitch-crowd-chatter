package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

type Word struct {
	Text      string
	Frequency int
}

func getPopularWords(messages []string) []Word {
	popularWords := make(map[string]int)

	for _, message := range messages {
		messageWords := strings.Split(message, " ")
		// make messageWords unique
		uniqueWords := make(map[string]bool)
		for _, word := range messageWords {
			uniqueWords[word] = true
		}

		for word := range uniqueWords {
			popularWords[word]++
		}
	}

	delete(popularWords, "")

	var words []Word
	for text, frequency := range popularWords {
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
	var botName = os.Args[1]
	var botToken = os.Args[2]
	var channelName = os.Args[3]
	client := twitch.NewClient(botName, botToken)

	// store the last 50 messages and remove the oldest one if there are more than 50 messages
	var messages []string

	var lastMessageSent = ""

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		messages = append(messages, message.Message)
		if len(messages) > 50 {
			messages = messages[1:]
		}

		popularWords := getPopularWords(messages)

		mostPopularWord := popularWords[0]
		fmt.Println(mostPopularWord)

		if mostPopularWord.Frequency >= 20 && mostPopularWord.Text != lastMessageSent {
			fmt.Println("Sending message:", mostPopularWord.Text)
			client.Say(channelName, mostPopularWord.Text)
			lastMessageSent = mostPopularWord.Text
		}

	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
