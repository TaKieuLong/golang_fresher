package controller

import (
	"fmt"

	"github.com/TaKieuLong/golang_fresher/internal/service"
	"github.com/TaKieuLong/golang_fresher/internal/vo"
	"github.com/TaKieuLong/golang_fresher/pkg/response"
	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService  *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserSerivce(),
// 	}
// }
// // controller->service->repo->models->dbs
// func (uc *UserController) GetUserByID(c *gin.Context){

// 	response.SuccessResponse(c,20001, uc.userService.GetInfoUser())
// }

type UserController struct {
	userService service.IUserService 
}


func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	params := vo.UserRegisterRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrParamInvalid, err.Error())
		return
	}
	fmt.Printf("Email param:: %s", params.Email)	
	result := uc.userService.Register(params.Email, params.Purpose)
response.SuccessResponse(c,20001,result)
}