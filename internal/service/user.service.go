package service

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/TaKieuLong/golang_fresher/internal/repo"
	"github.com/TaKieuLong/golang_fresher/internal/utils/random"
	"github.com/TaKieuLong/golang_fresher/internal/utils/sendto"

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
	hash := sha256.Sum256([]byte(email))
	hashEmail := fmt.Sprintf("%x", hash)
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
	err = sendto.SendTemplateEmailOtp([]string{email}, email, "otp.auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	if err != nil {	
		return response.ErrInvalidOTP
	}
	//
	return response.ErrCodeSuccess									
}

