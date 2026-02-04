package initialize

import (
	"fmt"

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


func InitRouter() *gin.Engine{

r:= gin.Default()

return  r
}