package nosql

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var redisAddress = fmt.Sprintf("%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
var redisDb, _ = strconv.Atoi(os.Getenv("REDIS_DB"))

var redisClient = redis.NewClient(&redis.Options{
	Addr:     redisAddress,
	Password: "",
	DB:       redisDb,
})

//GetRedisClient provide an instace of redis
func GetRedisClient() *redis.Client {
	_, err := redisClient.Ping().Result()

	if err != nil {
		panic(err)
	}

	return redisClient
}
