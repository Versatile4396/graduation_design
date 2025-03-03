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
}
