package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MSSkowron/mscache/client"
)

func main() {
	SendStuff()
}

func SendStuff() {
	c, err := client.New(":3000")
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		var (
			key   = []byte(fmt.Sprintf("key_%d", i))
			value = []byte(fmt.Sprintf("value_%d", i))
		)

		if err := c.Set(context.TODO(), key, value, 1000000000000000); err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
	}
}
