package initialize

import (
	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/TaKieuLong/golang_fresher/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}