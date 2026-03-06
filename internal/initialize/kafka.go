package initialize

import (
	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// Init Kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr: kafka.TCP("localhost:19092"),
		Topic: "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		global.Logger.Error("Failed to close Kafka producer", zap.Error(err))
	}
}
