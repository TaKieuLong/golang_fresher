package user

import (
	"github.com/TaKieuLong/golang_fresher/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	userController,_ := wire.InitUserRouterHanlder()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
	
	}
}