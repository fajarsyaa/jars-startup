package service

import (
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UserRegister(request request.RegisterUserRequest) (*model.User, error)
	Login(request request.LoginRequest) (*model.User, error)
	CheckAvailableEmail(request request.AvailableEmailRequest) (bool, error)
	SaveAvatar(Id, filepath string) (*model.User, error)
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

	NewUser, err := us.usrRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return NewUser, nil
}

func (us *userService) Login(request request.LoginRequest) (*model.User, error) {
	email := request.Email
	password := request.Password

	user, err := us.usrRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) CheckAvailableEmail(request request.AvailableEmailRequest) (bool, error) {
	email := request.Email
	user, err := us.usrRepo.FindUserByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID != "" {
		return false, nil
	}

	return true, nil
}

func (us *userService) SaveAvatar(Id, filepath string) (*model.User, error) {
	user, err := us.usrRepo.FindUserById(Id)
	if err != nil {
		return nil, err
	}

	user.AvatarFileName = filepath

	UpdatedUser, err := us.usrRepo.Update(*user)
	if err != nil {
		return nil, err
	}

	return UpdatedUser, nil
}
