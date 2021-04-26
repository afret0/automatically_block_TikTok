package source

import "github.com/go-redis/redis"

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     Config.GetString("redis.addr"),
		Password: Config.GetString("redis.password"),
		DB:       Config.GetInt("redis.DB"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		Logger.Fatalln(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
