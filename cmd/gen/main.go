package main

import (
	"github.com/TaKieuLong/golang_fresher/internal/initialize"
)

func main() {
	initialize.InitMySql() // connect DB
	initialize.GenTableDAO()   // generate model
}
