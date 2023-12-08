package application

import (
	"GinBoilerplate/config"
	"GinBoilerplate/infrastructure/database"
	redis_client "GinBoilerplate/infrastructure/redis"
	userController "GinBoilerplate/internal/domain/user/controller"
	userRepository "GinBoilerplate/internal/domain/user/repository"
	userService "GinBoilerplate/internal/domain/user/service"
	"GinBoilerplate/middleware/error"
	"GinBoilerplate/middleware/log"
	"GinBoilerplate/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Start() {
	app := gin.Default()
	app.Use(log.Logger())
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	app.Use(func(c *gin.Context) {
		error.ErrorHanlder(c)
	})
	cfg := config.Config
	fmt.Println("Config : ", cfg)
	db := database.ConnectDatabase(cfg)
	clientRedis := redis_client.ConnectRedis(cfg)

	err := error.LoadErrorListFromJsonFile(cfg.ErrorContract.JSONPathFile)
	if err != nil {
		logrus.Fatal(err)
	}

	// Define Repositories
	redisRepo := redis_client.NewRedisRepository(clientRedis)
	userRepo := userRepository.NewRepository()

	//Todo : Define Service here
	userSvc := userService.NewService(db.DB, userRepo, &redisRepo)

	//Todo: Define controller
	userCtrl := userController.NewController(userSvc)

	routerApp := router.NewRouter(&router.RouteParams{
		userCtrl,
	})
	routerApp.SetupRoute(app)

	address := cfg.AppConfig.Host + ":" + cfg.AppConfig.Port
	app.Run(address)
}
