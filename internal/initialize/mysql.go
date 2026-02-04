package initialize

import (
	"fmt"
	"time"

	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/TaKieuLong/golang_fresher/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))

		panic(err)
	}
}
func InitMySql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("InitMysql success")
	global.Mdb = db
	SetPool()
	GenTableDAO()
	// migrateTables()
}

// InitMySql().SetPool()
func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConnect)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func GenTableDAO(){
	  g := gen.NewGenerator(gen.Config{
    OutPath: "./internal/model",
    Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

  // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
  g.UseDB(global.Mdb)
  g.GenerateModel("go_crm_user") // reuse your gorm db

//   // Generate basic type-safe DAO API for struct `model.User` following conventions
//   g.ApplyBasic(model.User{})

//   // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

  // Generate the code
  g.Execute()
}
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migrating table error:", err)
	}
}
