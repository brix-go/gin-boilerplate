package database

import (
	"GinBoilerplate/config"
	"GinBoilerplate/internal/domain/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type Database struct {
	DB *gorm.DB
}

var dbInstance *Database
var once sync.Once

func newPostgres(cfg config.ConfigModel, gormConfig gorm.Config) *gorm.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Database)
	fmt.Println("DSN : ", dsn)

	//TODO : Change this driver if you use beside postgres (ex: mysql.Open())
	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}

func newMysql(cfg config.ConfigModel, gormConfig gorm.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	if cfg.Database.Password == "" {
		dsn = fmt.Sprintf("%s@tcp(%s:%d)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	}
	fmt.Println("DSN : ", dsn)

	//TODO : Change this driver if you use beside postgres (ex: mysql.Open())
	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}

func ConnectDatabase(cfg config.ConfigModel) *Database {
	var db *gorm.DB
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dbConfig := gorm.Config{
		Logger: dbLogger,
	}
	switch cfg.Database.Driver {
	case "mysql":
		db = newMysql(cfg, dbConfig)
	case "postgres":
		db = newPostgres(cfg, dbConfig)
	}
	fmt.Println("Database Connet Successfully")

	err := db.AutoMigrate(
		user.User{},
	)
	if err != nil {
		panic("Failed to migrate model to the database : " + err.Error())
	}
	dbInstance = &Database{DB: db}
	return dbInstance
}
