package main

import (
	"fmt"
	"strconv"
	"time"
)

func producer(s string, ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- s + strconv.Itoa(i)
		fmt.Println("producer: " + s + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

func consumer(ch <-chan string) {
	for v := range ch {
		fmt.Println("consumer: " + v)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan string)
	defer func() {
		close(ch)
	}()
	go producer("A", ch)
	go producer("B", ch)
	go producer("C", ch)
	go consumer(ch)

	time.Sleep(15 * time.Second)
}
