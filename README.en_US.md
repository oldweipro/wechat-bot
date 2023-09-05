# wechat-bot

[简体中文](README.md) | English

# OpenWeChat GPT-3.5 Turbo Chatbot

This is a simple Go program that demonstrates how to create a chatbot using the OpenWeChat framework and OpenAI's GPT-3.5 Turbo model.

## Prerequisites

Before running this code, make sure you have the following:

- Go programming language installed on your system.
- An OpenAI API key (`sk-xxx`) for accessing the GPT-3.5 Turbo model.
- You can also register for free at https://chat.oldwei.com to obtain an API key (`sk-xxx`).

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/oldweipro/wechat-bot.git
```

2. Compile and run the Go program:

```bash
cd yourrepository
go run main.go
```

## Usage

This program sets up a WeChat bot that listens for incoming messages and responds using the GPT-3.5 Turbo model from OpenAI. It performs the following actions:

- Initializes a WeChat bot in desktop mode.
- Handles incoming text messages.
- Replies to messages either in a group if mentioned or directly in a one-on-one chat.
- Utilizes the GPT-3.5 Turbo model for chat completions.

## Configuration

Ensure you have replaced `"sk-xxx"` with your actual OpenAI API key in the `ChatCompletion` function.

```go
config := openai.DefaultConfig("sk-xxx")
```

## License

This code is provided under the [MIT License](LICENSE).

## Author

- 老卫同学
- GitHub: [oldweipro](https://github.com/oldweipro)
