package user

import (
	"GinBoilerplate/internal/domain/user/dto/requests"
	"GinBoilerplate/internal/domain/user/dto/responses"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	Email      string
	Password   string
	Phone      string
	VerifiedAt sql.NullTime `gorm:"default:null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func NewUser(request requests.RegisterRequest) User {
	return User{
		ID:       uuid.NewString(),
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}
}

type UserRepository interface {
	FindUserByEmail(db *gorm.DB, email string) (*User, error)
	CreateUser(db *gorm.DB, user *User) (*User, error)
	UpdateUser(db *gorm.DB, updatedUser *User) (*User, error)
	GetUserByID(db *gorm.DB, id string) (*User, error)
	GetAllUser(db *gorm.DB) ([]*User, error)
	DeleteUser(db *gorm.DB, id string) error
	VerifyUser(db *gorm.DB, id string) error
}

type UserService interface {
	Login(ctx context.Context, userRequest *requests.LoginRequest) (res *responses.LoginResponse, err error)
	RegisterUser(ctx context.Context, userRequest *requests.RegisterRequest) (user User, err error)
	GetDetailUserById(ctx context.Context, id string) (user responses.UserDetail, err error)
	VerifyUser(ctx context.Context, verReq requests.VerifiedUserRequest) error
	ResendOTP(ctx context.Context, userId string) error
}

type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetDetailUserJWT(ctx *gin.Context)
	VerifyUser(ctx *gin.Context)
	ResendOTP(ctx *gin.Context)
}
