package service

import (
	"github.com/hoangxt/go-ecommerce-ba-api/internal/repo"
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/response"
)

// type UserService struct {
// 	UserRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		UserRepo: repo.NewUserRepo(),
// 	}
// }

// // us user service
// func (us *UserService) GetInfoUser() string {
// 	return us.UserRepo.GetInfoUser()
// }

// Interface

type IUserService interface {
	Register(email string, purpose string) int
	// ...
}

type userService struct {
	userRepo repo.IUserRepository
	// ...
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// Check if email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
