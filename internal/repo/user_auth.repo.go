package repo

import (
	"fmt"
	"time"

	"github.com/TaKieuLong/golang_fresher/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct {
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

func (ur *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	// ctx is defined in common.repo.go
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)*time.Second).Err()
}
