package main

import "fmt"
import "github.com/go-redis/redis"
func main(){
	client := redis.NewClient(&redis.Options{
		Addr:     "47.88.222.158:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	a:=client.Get("a")
	l:=client.LIndex("l",0)
	fmt.Println(pong, err)
	fmt.Println(a.String())
	fmt.Println(l.String())
}
