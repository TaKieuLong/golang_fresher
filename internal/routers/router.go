package routers

import (
	"github.com/TaKieuLong/golang_fresher/internal/controller"
	"github.com/gin-gonic/gin"
)
	

func NewRouter() *gin.Engine{
r:= gin.Default()
v1:=r.Group("/v1/2024")
{
	v1.GET("/user/1",controller.NewUserController().GetUserByID)
	
}
return  r
}