package repository

import (
	"GinBoilerplate/internal/domain/user"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
}

func NewRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindUserByEmail(db *gorm.DB, email string) (user *user.User, err error) {
	err = db.Debug().Take(&user, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(db *gorm.DB, user *user.User) (*user.User, error) {
	err := db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(db *gorm.DB, updatedUser *user.User) (*user.User, error) {
	return updatedUser, nil
}

func (r *userRepository) GetUserByID(db *gorm.DB, id string) (user *user.User, err error) {
	err = db.Debug().First(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllUser(db *gorm.DB) (user []*user.User, err error) {
	err = db.Debug().Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(db *gorm.DB, id string) error {
	return nil
}

func (r *userRepository) VerifyUser(db *gorm.DB, id string) error {
	err := db.Debug().Model(&user.User{}).Where("id = ?", id).Update("verified_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
