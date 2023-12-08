package redis_client

import (
	"GinBoilerplate/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.ConfigModel) *redis.Client {
	fmt.Println("REDIS HOST : ")
	//TODO: Add your redis address
	store := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port), Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", DB: 0})
	fmt.Println("Redis client connected")
	return store
}
