package init

import (
	configuratio "codeCademy/configuration"
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)
var RedisClient *redis.Client


func InitRedis(){

  config := configuratio.GlobalConfiguration
  RedisClient = redis.NewClient(&redis.Options{
    Addr: config.Redis.Address,
    Password: config.Redis.Password,
    DB:0,
  })

  ping, err := RedisClient.Ping(context.Background()).Result()
  if err !=nil{
    panic("Error connecting to redis server! - "+ err.Error())
  }
  
  fmt.Println(ping)
}
