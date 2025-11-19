package global

import (
	"github.com/TaKieuLong/golang_fresher/pkg/logger"
	"github.com/TaKieuLong/golang_fresher/pkg/setting"
	"gorm.io/gorm"
)
var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB

)