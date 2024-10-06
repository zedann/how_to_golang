package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Response struct {
	val int
	err error
}

func main() {

	start := time.Now()
	ctx := context.Background()

	userID := 1
	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))

}

func fetchUserData(ctx context.Context, userID int) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*150)
	defer cancel()

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdParty()

		respch <- Response{
			val: val,
			err: err,
		}
	}()

	for {
		select {
		case res := <-respch:
			return res.val, res.err
		case <-ctx.Done():
			return 0, fmt.Errorf("third party took too long")
		}
	}

}

func fetchThirdParty() (int, error) {
	// inconsistant could take 500ms / 2000ms / 20ms we do not know
	max := 500
	min := 20
	time.Sleep(time.Millisecond * (time.Duration(rand.Intn(max-min+1) + min)))

	return 666, nil

}
