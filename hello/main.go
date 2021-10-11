package main

import (
    "context"
    "fmt"
    "github.com/themarcelor/helloworld/redisstuff"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {
    fmt.Println(helloworld())
    ctx := context.TODO()

    redis := redisstuff.GetRedisClient()

    redis.Set(ctx, "language", "Go", 0)
    language := redis.Get(ctx, "language")
    year := redis.Get(ctx, "year")

    fmt.Println(language.Val()) // "Go"
    fmt.Println(year.Val()) // ""
}
