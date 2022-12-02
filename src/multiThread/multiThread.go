package multiThread

import (
	"fmt"
	"time"
)

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}

func MultiThread() {
	channel := make(chan int)

	// go func() {
	for i := 0; i <= 50; i = i + 1 {
		go worker(channel)
	}
	// }()

	for i := 0; i <= 100; i = i + 1 {
		channel <- i
	}
}
