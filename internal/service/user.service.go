package service

import (
	"fmt"

	"github.com/TaKieuLong/golang_fresher/internal/repo"
	"github.com/TaKieuLong/golang_fresher/pkg/crypto"
	"github.com/TaKieuLong/golang_fresher/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
userRepo repo.IUserRepository,
userAuthRepo repo.IUserAuthRepository) IUserService{
	return &userService{
		userRepo: userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	hashEmail := crypto.GetHash(email)
	fmt.Sprintf("hashEmail: %s", hashEmail)
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	otp := random.GenerateSixDigitOtp()

	if purpose == "TEST_USER" {
		otp = 123456
	}
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	//semdEmail OTP
	err = sendto.SendTextEmailOtp([]string{email}, email, otp)
	if err != nil {
		return response.ErrInvalidOTP
	}
	//
	return response.ErrCodeSuccess									
}

