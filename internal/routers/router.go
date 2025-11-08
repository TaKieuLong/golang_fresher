package routers

import (
	"fmt"

	"github.com/TaKieuLong/golang_fresher/internal/controller"
	"github.com/TaKieuLong/golang_fresher/internal/middlewares"
	"github.com/gin-gonic/gin"
)
	
func AA() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("Before--->AA")
		c.Next()
		fmt.Println("Alter--->AA")
	}
}

func BB() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("Before--->BB")
		c.Next()
		fmt.Println("Alter--->BB")
	}
}

func CC(c *gin.Context){
		fmt.Println("Before--->CC")
		c.Next()
		fmt.Println("Alter--->CC")
	}


func NewRouter() *gin.Engine{

r:= gin.Default()
r.Use(middlewares.AuthenMiddleware(),BB(),CC)
v1:=r.Group("/v1/2024")
{
	v1.GET("/user/1",controller.NewUserController().GetUserByID)
	
}
return  r
}