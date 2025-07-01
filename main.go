package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

		// 记录开始时间
		start := time.Now()

		// 打印接收到的消息
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// 计算处理用时
		elapsed := time.Since(start)

		// 回复消息，包含处理用时
		responseText := fmt.Sprintf("Golang 你说：%s\n处理用时：%v", update.Message.Text, elapsed)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)

		// 发送消息
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}

		// 记录总用时
		totalElapsed := time.Since(start)
		log.Printf("消息处理完成 - 用户: [%s], 消息: [%s], 总用时: %v",
			update.Message.From.UserName,
			update.Message.Text,
			totalElapsed)
	}
}
