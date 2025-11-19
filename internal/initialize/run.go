package initialize

import (
	"fmt"

	"github.com/TaKieuLong/golang_fresher/global"
	"go.uber.org/zap"
)

func Run() {
	Loadconfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok!",zap.String("ok","success"))
	InitMySql()
	InitRedis()
	InitRouter()
	r := InitRouter()

	r.Run(":8002")
}