package controller

import (
	"github.com/TaKieuLong/golang_fresher/internal/service"
	"github.com/TaKieuLong/golang_fresher/pkg/response"
	"github.com/gin-gonic/gin"
)
type UserController struct {
	userService  *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserSerivce(),
	}
}
// controller->service->repo->models->dbs
func (uc *UserController) GetUserByID(c *gin.Context){
	
	response.SuccessResponse(c,20001, uc.userService.GetInfoUser())
}