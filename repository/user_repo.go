package repository

import (
	"bwu-startup/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserById(id string) (*model.User, error)
	Update(user model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *model.User) (*model.User, error) {
	err := ur.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := ur.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) FindUserById(id string) (*model.User, error) {
	user := model.User{}
	err := ur.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(user model.User) (*model.User, error) {
	err := ur.db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
