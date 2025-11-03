package service

import "github.com/TaKieuLong/golang_fresher/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserSerivce() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}