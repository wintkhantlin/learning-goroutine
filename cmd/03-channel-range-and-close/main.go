package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Int()
		time.Sleep(200000000)
		c <- rand.Int()
	}
	close(c)
}

func main() {
	c := make(chan int)

	go process(c)

	for i := range c {
		fmt.Println(i)
	}
}
