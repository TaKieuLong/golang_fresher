package main

import (
	"github.com/TaKieuLong/golang_fresher/internal/routers"
)

func main() {
  r := routers.NewRouter()
 
  r.Run(":8002")
}