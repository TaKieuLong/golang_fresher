package repo

import "github.com/TaKieuLong/golang_fresher/internal/model"

type IUserAuthRepository interface {
AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository() IUserAuthRepository {}

func (ur *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	
	key:= fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}