# Twitch Crowd Chatter

Twitch Crowd Chatter is a Go application designed to interact with Twitch chat in various ways, such as collecting chat messages, analyzing popular words, and responding to chat with automated messages based on the chat's content. This project leverages the github.com/gempir/go-twitch-irc/v4 library to connect to Twitch IRC for real-time chat interaction.

## Requirements

- Go version 1.21.0 or higher
- A Twitch OAuth token for authentication

## Installation

1. Ensure you have Go installed on your system. You can download it from the official Go website.
2. Clone the repository to your local machine:

```bash
git clone https://github.com/lebogoo/twitch-crowd-chatter.git
```

3. Navigate to the project directory:

```bash
cd twitch-crowd-chatter
```

4. Install the required dependencies:

```bash
go mod tidy
```

## Usage

The application supports multiple modes of operation, each triggered by a different subcommand:

### Collecting Chat Messages

To collect chat messages from a specific channel and save them to a file:

```bash
./twitch-crowd-chatter collect <channelName>
```

This will create or append to a file named `<channelName>.txt`, storing the timestamp, username, and message content.

### Processing Chat Messages

To process previously collected chat messages and analyze the top chatters and other statistics, run:

```bash
./twitch-crowd-chatter process <channelName>
```

### Analyzing and Responding to Chat

To analyze chat messages for popular words and potentially respond based on predefined criteria:

```bash
./twitch-crowd-chatter chatter <botName> <botToken> <channelName>
```

This mode requires a Twitch OAuth token for the bot to authenticate and send messages.

## Modules

The application is structured into several modules, each responsible for a different aspect of the application's functionality:

- Chatter: Handles analyzing chat messages and responding to chat.
- Collect: Collects chat messages and saves them to a file.
- Process: Processes and sorts data, such as chat messages. This keeps the original data intact.

# Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to discuss potential improvements or features.

# License

This project is licensed under the MIT License - see the LICENSE file for details.
