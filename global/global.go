package global

import (
	"github.com/TaKieuLong/golang_fresher/pkg/logger"
	"github.com/TaKieuLong/golang_fresher/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
