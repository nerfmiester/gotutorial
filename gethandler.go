// GetHandler
package main

import (
	"fmt"

	"github.com/ivpusic/neo"
	"gopkg.in/redis.v2"
)

func GetHandler(ctx *neo.Ctx) {
	//var role bubbleRole
	client := redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	keys := client.Keys("*")

	for _, value := range keys.Val() {

		if (client.TTL(value).Val()) < 0 {
			fmt.Printf("value = %d", value)
		}

	}
}
