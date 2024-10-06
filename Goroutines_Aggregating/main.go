package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	userName := fetchUser()

	resCh := make(chan any, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikes(userName, resCh, wg)
	go fetchUserMatch(userName, resCh, wg)

	wg.Wait() // block until we had 2 wg.Done() calls
	close(resCh)

	for res := range resCh {
		fmt.Println("response", res)
	}

	fmt.Println("time:", time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, resCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	resCh <- 11
	wg.Done()

}

func fetchUserMatch(userName string, resCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	resCh <- "Anna"
	wg.Done()
}
