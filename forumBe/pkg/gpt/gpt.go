package gpt

import (
	"forum/global"
	"os"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
)

func Init() {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
		arkruntime.WithTimeout(2*time.Minute),
		arkruntime.WithRetryTimes(2),
	)
	global.Gpt = client
	// ctx := context.Background()

	// fmt.Println("----- standard request -----")
	// const reqstr = "请总结一下下面这段内容："
	// req := model.CreateChatCompletionRequest{
	// 	Model: "doubao-1-5-vision-pro-32k-250115",
	// 	Messages: []*model.ChatCompletionMessage{
	// 		{
	// 			Role: model.ChatMessageRoleSystem,
	// 			Content: &model.ChatCompletionMessageContent{
	// 				StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
	// 			},
	// 		},
	// 		{
	// 			Role: model.ChatMessageRoleUser,
	// 			Content: &model.ChatCompletionMessageContent{
	// 				StringValue: volcengine.String("请总结一下下面这段内容：" + reqstr),
	// 			},
	// 		},
	// 	},
	// }

	// resp, err := client.CreateChatCompletion(ctx, req)
	// if err != nil {
	// 	fmt.Printf("standard chat error: %v\n", err)
	// 	return
	// }
	// fmt.Println(*resp.Choices[0].Message.Content.StringValue)
}
