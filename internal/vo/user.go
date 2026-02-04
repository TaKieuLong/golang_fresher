package vo

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" `
	Purpose  string `json:"purpose"`
}
