package service

import (
	"bwu-startup/helper/jwt_token"
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UserRegister(request request.RegisterUserRequest) (*model.User, *string, error)
	Login(request request.LoginRequest) (*model.User, *string, error)
	CheckAvailableEmail(request request.AvailableEmailRequest) (bool, error)
	SaveAvatar(Id, filepath string) (*model.User, error)
	GetUserById(Id string) (*model.User, error)
}

type userService struct {
	usrRepo  repository.UserRepository
	jwtToken jwt_token.JwtToken
}

func NewUserService(userRepo repository.UserRepository, jwtToken jwt_token.JwtToken) *userService {
	return &userService{usrRepo: userRepo, jwtToken: jwtToken}
}

func (us *userService) UserRegister(request request.RegisterUserRequest) (*model.User, *string, error) {
	user := model.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}
	user.PasswordHash = string(passwordhash)
	user.ID = uuid.New().String()
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	token, err := us.jwtToken.GenerateToken(user.ID)
	if err != nil {
		return nil, nil, err
	}

	NewUser, err := us.usrRepo.Create(&user)
	if err != nil {
		return nil, nil, err
	}

	return NewUser, token, nil
}

func (us *userService) Login(request request.LoginRequest) (*model.User, *string, error) {
	email := request.Email
	password := request.Password

	logginedUser, err := us.usrRepo.FindUserByEmail(email)
	if err != nil {
		return nil, nil, err
	}
	if logginedUser.ID == "" {
		return nil, nil, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(logginedUser.PasswordHash), []byte(password))
	if err != nil {
		return nil, nil, err
	}

	token, err := us.jwtToken.GenerateToken(logginedUser.ID)
	if err != nil {
		return nil, nil, err
	}

	return logginedUser, token, nil
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

func (us *userService) GetUserById(Id string) (*model.User, error) {
	user, err := us.usrRepo.FindUserById(Id)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user with that id is not found")
	}

	return user, nil
}
