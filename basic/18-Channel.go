package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender1(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "sender1"
}

func sender2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "sender2"
}

func main() {
	// https://colobu.com/2016/04/14/Golang-Channels/
	// declare two unbuffered channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// select
	go sender1(ch1)
	go sender2(ch2)

	// like .join() in java
	// main thread will be blocked until receiving channel's vaule
	s, t := <-ch1, <-ch2

	fmt.Println(s, t)

	// use range to get the uncertain length of values
	n := make(chan int)
	go func(end int, n chan int) {
		for i := 0; i < end; i++ {
			n <- i
		}
		close(n)
	}(rand.Int()%10, n)

	fmt.Println("Start receiving")
	for v := range n {
		fmt.Printf("%d ", v)
	}

	// Buffered channel do not cause block
	ch3 := make(chan int, 8)
	go func() {
		ch3 <- 8
		time.Sleep(1 * time.Second)
		fmt.Println("buffer channel")
	}()
	<-ch3
}
