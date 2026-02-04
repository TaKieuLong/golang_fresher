package repo

import (
	"fmt"
	"time"

	"github.com/TaKieuLong/golang_fresher/global"
	"gorm.io/gorm"
)

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