package service

import (
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UserRegister(request request.RegisterUserRequest) (*model.User, error)
}

type userService struct {
	usrRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{usrRepo: userRepo}
}

func (us *userService) UserRegister(request request.RegisterUserRequest) (*model.User, error) {
	user := model.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(passwordhash)
	user.ID = uuid.New().String()
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	NewUser, err := us.usrRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	return NewUser, nil
}
