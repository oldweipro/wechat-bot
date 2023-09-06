package main

import (
	"context"
	"github.com/eatmoreapple/openwechat"
	"github.com/sashabaranov/go-openai"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		if (msg.IsSendByGroup() && msg.IsAt()) || !msg.IsSendByGroup() {
			gptResp := ChatCompletion(msg.Content)
			msg.ReplyText(gptResp)
		}
		msg.AsRead()
	}) // 只处理消息类型为文本类型的消息
	bot.MessageHandler = dispatcher.AsMessageHandler()                  // 注册消息处理函数
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl                      // 注册登陆二维码回调
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json") // 创建热存储容器对象
	defer reloadStorage.Close()
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()) // 执行热登录，热登录有一点缺点就是它的有效期很短（具体多久我也不知道），但是好像一直有消息收发就不会过期
	bot.Block()                                                   // 阻塞主goroutine, 直到发生异常或者用户主动退出
}
func ChatCompletion(content string) string {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo0301,
		MaxTokens: 1000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	}
	//您可以在 https://chat.oldwei.com 免费注册以获取 API 密钥。
	config := openai.DefaultConfig("sk-xxx")
	config.BaseURL = "https://api.chat.oldwei.com/v1"
	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "本宝宝业务繁忙，请稍后重试！"
	}
	return response.Choices[0].Message.Content
}
