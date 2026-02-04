package repo

import (
	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/TaKieuLong/golang_fresher/internal/model"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "Longi"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}
type userRepository struct {}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	//SELECT * FROM users WHERE email = email order by id desc limit 1
	row := global.Mdb.Table(TableUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull	
}