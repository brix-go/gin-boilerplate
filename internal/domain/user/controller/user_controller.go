package controller

import (
	"GinBoilerplate/internal/domain/user"
	"GinBoilerplate/internal/domain/user/dto/requests"
	middleware "GinBoilerplate/middleware/error"
	validation "GinBoilerplate/middleware/validate"
	"GinBoilerplate/shared"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type userController struct {
	userService user.UserService
}

func NewController(service user.UserService) user.UserController {
	return &userController{
		userService: service,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var loginReq requests.LoginRequest
	err := ctx.ShouldBind(&loginReq)
	if err != nil {
		ctx.Error(errors.New(shared.ErrInvalidRequestFamily))
		ctx.Abort()
		return
	}
	fieldErr := validation.ValidateRequest(ctx, loginReq)
	if fieldErr != nil {
		ctx.Abort()
		return
	}
	res, err := c.userService.Login(ctx, &loginReq)
	if err != nil {
		ctx.Error(errors.New(shared.ErrUnexpectedError))
		ctx.Abort()
		return
	}
	middleware.ResponseSuccess(ctx, shared.RespSuccess, res)
	return
}

func (c *userController) Register(ctx *gin.Context) {
	var registerRequest requests.RegisterRequest
	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		ctx.Error(errors.New(shared.ErrInvalidRequestFamily))
		ctx.Abort()
		return
	}
	fieldErr := validation.ValidateRequest(ctx, registerRequest)
	if fieldErr != nil {
		ctx.Abort()
		return
	}
	_, err = c.userService.RegisterUser(ctx, &registerRequest)
	if err != nil {
		ctx.Error(errors.New(shared.ErrUnexpectedError))
		ctx.Abort()
		return
	}
	middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
	return
}

func (c *userController) GetDetailUserJWT(ctx *gin.Context) {
	data, err := c.userService.GetDetailUserById(ctx, "1")
	if err != nil {
		ctx.Error(errors.New(shared.ErrInvalidRequestFamily))
		ctx.Abort()
		return
	}
	middleware.ResponseSuccess(ctx, shared.RespSuccess, data)
	return
}

func (c *userController) VerifyUser(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
	return
}

func (c *userController) ResendOTP(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
	return
}
