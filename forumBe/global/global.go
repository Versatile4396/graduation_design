package global

import (
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Gpt *arkruntime.Client
	// Snowflake *snowflake.Snowflake
)
