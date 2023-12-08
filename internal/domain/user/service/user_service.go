package service

import (
	redis_client "GinBoilerplate/infrastructure/redis"
	"GinBoilerplate/internal/domain/user"
	"GinBoilerplate/internal/domain/user/dto/requests"
	"GinBoilerplate/internal/domain/user/dto/responses"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type userService struct {
	db        *gorm.DB
	repo      user.UserRepository
	redisRepo *redis_client.Repository
}

func NewService(db *gorm.DB, repo user.UserRepository, redis *redis_client.Repository) user.UserService {
	return &userService{
		db:        db,
		repo:      repo,
		redisRepo: redis,
	}
}

func (s *userService) Login(ctx context.Context, userRequest *requests.LoginRequest) (*responses.LoginResponse, error) {
	//TODO: Implement login
	return &responses.LoginResponse{}, nil
}

func (s *userService) RegisterUser(ctx context.Context, userRequest *requests.RegisterRequest) (userData user.User, err error) {
	tx := s.db.WithContext(ctx).Begin()
	//TODO: Implement register
	userData = user.NewUser(*userRequest)
	_, err = s.repo.CreateUser(tx, &userData)
	if err != nil {
		tx.Rollback()
		return userData, err
	}
	tx.Commit()
	return userData, nil

}

func (s *userService) GetDetailUserById(ctx context.Context, id string) (res responses.UserDetail, err error) {
	//TODO: Implement getDetailUserbyId
	tx := s.db.Begin()
	users, err := s.repo.GetAllUser(tx)
	if err != nil {
		return res, err
	}
	fmt.Println("User : ", users)
	return res, nil
}

func (s *userService) VerifyUser(ctx context.Context, verReq requests.VerifiedUserRequest) error {
	// TODO: Implement verify user
	return nil
}

func (s *userService) ResendOTP(ctx context.Context, userId string) error {
	return nil
}
