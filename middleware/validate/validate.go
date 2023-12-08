package middleware

import (
	"GinBoilerplate/shared"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type ErrorResponse struct {
	Error       string
	FailedField string
	Tag         string
	Value       interface{}
}

func ValidateRequest(ctx *gin.Context, data interface{}) error {
	validate := validator.New()
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				ctx.Error(errors.New(shared.ErrInvalidFieldFormat)).SetMeta(err.Field())
				return errors.New(shared.ErrInvalidFieldFormat)
			case "email":
				ctx.Error(errors.New(shared.ErrInvalidFieldFormat)).SetMeta(err.Field())
				return errors.New(shared.ErrInvalidFieldFormat)
			case "min":
				ctx.Error(errors.New(shared.ErrInvalidMinFieldFormat)).SetMeta(err.Field())
				return errors.New(shared.ErrInvalidMinFieldFormat)
			}
		}
	}

	return nil
}
