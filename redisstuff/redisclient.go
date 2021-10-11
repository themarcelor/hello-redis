package redisstuff

import (
   "github.com/go-redis/redis/v8"
)

func GetRedisClient()(*redis.Client) {
   client := redis.NewClient(&redis.Options{
      Addr:     "localhost:6379", // host:port of the redis server
      Password: "", // no password set
      DB:       0,  // use default DB
   })
   return client
}
