package repo

import (
	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/TaKieuLong/golang_fresher/internal/database"
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
type userRepository struct {
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	//SELECT * FROM users WHERE email = email order by id desc limit 1
	// row := global.Mdb.Table(TableUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	// return row != NumberNull	

	user, err := ur.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}