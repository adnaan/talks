package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//INIT OMIT
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	defer c.Close()

	//set
	c.Do("SET", "hello", "world")

	//get
	world, err := redis.String(c.Do("GET", "hello"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT

}
