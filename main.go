package main

import (
	"fmt"
	"github.com/thiagoaugustus/GoLang-Channels/goroutines"
)

func main() {
	c := make(chan int)
	go goroutines.Simple(c)

	value := <-c
	fmt.Println(value)
}
