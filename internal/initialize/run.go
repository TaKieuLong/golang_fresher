package initialize

func Run() {
	Loadconfig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitRouter()
	r := InitRouter()

	r.Run(":8002")
}