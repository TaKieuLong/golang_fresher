package initialize

import (
	"fmt"
	"time"

	"github.com/TaKieuLong/golang_fresher/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func checkErrorPanic(err error, errString string){
	if err != nil{
		global.Logger.Error(errString,zap.Error(err))
		panic(err)
	}
}
func InitMySql() {
	m := global.Config.Mysql
	 dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
var s = fmt.Sprintf(dsn,m.Username,m.Password,m.Host,m.Port,m.Dbname )
  db, err := gorm.Open(mysql.Open(s), &gorm.Config{
	SkipDefaultTransaction: false,
  })
checkErrorPanic(err, "InitMysql initialization error")
global.Logger.Info("InitMysql success")
global.Mdb = db

SetPool()
}
// InitMySql().SetPool()
func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::",err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConnect)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func migrateTables() {

}