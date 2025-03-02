package domain

import (
	"fmt"
	"log"
	"time"
)

func baseSendAndReceive() {
	fmt.Println("Start BaseSendAndReceive")
	ch := make(chan int)
	go func(number int) {
		time.Sleep(time.Second * 2)
		ch <- number
		fmt.Printf("Send Success %v To Channel \n", number)
	}(5)
	go func() {
		number := <-ch
		fmt.Printf("Fetch Success %v From Channel \n", number)
	}()
	time.Sleep(time.Second * 4)
	fmt.Println("End BaseSendAndReceive")
}

func sendMultiData() {
	log.Printf(">>>>>>>>>>>Start SendMultiData<<<<<<<<<<<<<<<")
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			log.Printf("Send Success To Channel Value: %v", i)
		}
		close(ch)
	}()
	for v := range ch {
		log.Printf("Fetch Value From Channel: %v", v)
	}

	// for {
	// 	number, ok := <-ch
	// 	if !ok {
	// 		break // Thoát khỏi vòng lặp nếu channel đã đóng
	// 	}
	// 	log.Printf("Fetch Value From Channel: %v, %v", number, ok)
	// }

	// for i := 1; i <= 10; i++ {
	// 	number := <-ch
	// 	log.Printf("Fetch Value From Channel: %v", number)
	// }
	log.Printf(">>>>>>>>>>>End SendMultiData<<<<<<<<<<<<<<<")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func executeSum() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c
	log.Printf("%v %v %v", x, y, x+y)
}

func selectDemo() {
	channelFirst := make(chan int)
	channelSecond := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		channelFirst <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		channelSecond <- 2
	}()
	for index := 1; index <= 2; index++ {
		select {
		case valueFirst := <-channelFirst:
			log.Printf("Value Of First Channel: %v", valueFirst)
		case valueSecond := <-channelSecond:
			log.Printf("Value Of Second Channel: %v", valueSecond)
		}
	}
}

// // MUTEX

func RunConcurrency() {
	// baseSendAndReceive()
	// sendMultiData()
	// executeSum()
	// selectDemo()
}
