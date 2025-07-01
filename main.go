package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// 从环境变量获取 bot token
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置 TELEGRAM_BOT_TOKEN 环境变量")
	}

	// 创建 bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bot 已启动: @%s", bot.Self.UserName)

	// 设置 webhook
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("请设置 WEBHOOK_URL 环境变量")
	}

	wh, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	// 创建更新通道
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":8443", nil)

	// 处理消息
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// 打印接收到的消息
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// 回复 "Hello, World!"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Golang 你说："+update.Message.Text)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
