package main

import (
	"fmt"
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
	}(10, n)

	fmt.Println("Start receiving")
	for v := range n {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
	fmt.Println()

	// Unbuffered channel causes block
	ch3 := make(chan int)
	go func() {
		ch3 <- 1
		fmt.Println("ok1")
		ch3 <- 2
		fmt.Println("ok2")
	}()
	<-ch3
	time.Sleep(3 * time.Second)
	<-ch3

	time.Sleep(5 * time.Second)
	fmt.Println()
	/*
		ok1
		(wait 5 second)
		ok2
	*/

	// Buffered channel do not cause block
	// Even if main routing have not read channel
	// go routing will still send value to channel

	ch4 := make(chan int, 5)
	go func() {
		ch4 <- 3
		fmt.Println("ok3")
		ch4 <- 4
		fmt.Println("ok4")
	}()
	<-ch4
	time.Sleep(5 * time.Second)
	<-ch4

	/*
		ok3
		(without waiting)
		ok4
	*/
}
