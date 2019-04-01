package data

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var redisAddress = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
var redisDb, _ = strconv.Atoi(os.Getenv("REDIS_DB"))

var redisClient = redis.NewClient(&redis.Options{
	Addr: redisAddress,
	DB:   redisDb,
})
