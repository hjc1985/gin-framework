package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mqenergy-go/pkg/lib"
)

var (
	DB     *gorm.DB      // Mysql数据库
	Logger *lib.Logger   // 日志
	Redis  *redis.Client // redis连接池
)
