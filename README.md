# Telegram Hello World Bot

这是一个简单的 Telegram webhook 机器人示例，它会对收到的任何消息回复 "Hello, World!"。

## 准备工作

1. 首先需要在 Telegram 上创建一个机器人。通过 [@BotFather](https://t.me/botfather) 创建机器人并获取 API Token。

2. 准备一个具有 HTTPS 的公网域名（Telegram webhook 要求必须使用 HTTPS）。可以使用 ngrok 等工具进行本地测试。

## 环境变量设置

运行程序前需要设置以下环境变量：

```bash
export TELEGRAM_BOT_TOKEN="你的机器人Token"
export WEBHOOK_URL="https://你的域名/webhook"
```

## 运行程序

1. 安装依赖：
```bash
go mod tidy
```

2. 运行程序：
```bash
go run main.go
```

## 功能说明

- 机器人会监听 8443 端口
- 对收到的任何消息都会回复 "Hello, World!"
- 所有接收到的消息都会在控制台打印日志

## 注意事项

- 确保服务器的 8443 端口已开放
- webhook URL 必须使用 HTTPS
- 建议在生产环境中使用反向代理（如 Nginx）来处理 HTTPS 