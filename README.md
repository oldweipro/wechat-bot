# WeChat 机器人

简体中文 | [English](README.en_US.md)

# OpenWeChat GPT-3.5 Turbo 聊天机器人

这是一个简单的 Go 程序，演示了如何使用 OpenWeChat 框架和 OpenAI 的 GPT-3.5 Turbo 模型创建聊天机器人。

## 先决条件

在运行此代码之前，请确保您具备以下条件：

- 在您的系统上安装了 Go 编程语言。
- 拥有一个用于访问 GPT-3.5 Turbo 模型的 OpenAI API 密钥（`sk-xxx`）。
- 您还可以在 https://chat.oldwei.com 免费注册以获取 API 密钥（`sk-xxx`）。

## 安装

1. 将此存储库克隆到您的本地计算机：

```bash
git clone https://github.com/oldweipro/wechat-bot.git
```

2. 编译并运行 Go 程序：

```bash
cd yourrepository
go run main.go
```

## 使用

该程序设置了一个 WeChat 机器人，监听传入的消息并使用来自 OpenAI 的 GPT-3.5 Turbo 模型进行回复。它执行以下操作：

- 初始化一个桌面模式的 WeChat 机器人。
- 处理传入的文本消息。
- 根据需要在群组或一对一聊天中回复消息。
- 利用 GPT-3.5 Turbo 模型进行聊天完成。

## 配置

请确保您已在 `ChatCompletion` 函数中将 `"sk-xxx"` 替换为您的实际 OpenAI API 密钥。

```go
config := openai.DefaultConfig("sk-xxx")
```

## 许可证

此代码在 [MIT 许可证](LICENSE) 下提供。

## 作者

- 老卫同学
- GitHub：[oldweipro](https://github.com/oldweipro)