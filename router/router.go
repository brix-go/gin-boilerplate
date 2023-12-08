package router

import (
	"GinBoilerplate/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type RouteParams struct {
	//TODO: Define your controller
	user.UserController
}

type RouterStruct struct {
	RouteParams RouteParams
}

func NewRouter(params *RouteParams) RouterStruct {
	return RouterStruct{
		RouteParams: *params,
	}
}

func (r *RouterStruct) SetupRoute(app *gin.Engine) {
	v1 := app.Group("/api/v1")

	authRouter := v1.Group("/auth")
	authRouter.POST("/login", r.RouteParams.Login)
	authRouter.POST("/register", r.RouteParams.Register)

	userRouter := v1.Group("/users")
	userRouter.GET("/", r.RouteParams.GetDetailUserJWT)

}
