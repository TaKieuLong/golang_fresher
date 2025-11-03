package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
// sugar:= zap.NewExample().Sugar()
// sugar.Info("hello %s, age :%d", "long", 40)

// logger := zap.NewExample()
// logger.Info("Hello", zap.String("name","Long"))

//Để hiểu cách log từng cách
// logger := zap.NewExample()
// logger.Info("Heloo NewExample")
// //Development
// logger,_ = zap.NewDevelopment()
// logger.Info("Heloo NewDevelopment")

// //Producttion
// logger,_ = zap.NewProduction()
// logger.Info("Heloo NewProduction")


// custom
encoder := getEncoderLog()
sync := getWriterSync()
core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

logger := zap.New(core, zap.AddCaller())

logger.Info("Info log", zap.Int("line",1))
logger.Error("Error log", zap.Int("line",2))
}


func getEncoderLog() zapcore.Encoder{
	encodeConfig := zap.NewProductionEncoderConfig()
	//	1762091648.9400837 ---> 2025-11-02T20:54:08.939+0700
	 encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	 //	ts--->Time
	 encodeConfig.TimeKey ="Time"
	 // info -> INFO
	 encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	 encodeConfig.LevelKey = "information"

	 //	caller
	 encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	 return zapcore.NewJSONEncoder(encodeConfig)
}

//
func getWriterSync() zapcore.WriteSyncer {
	file,_ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole,syncFile)
}