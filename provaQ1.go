package main

import (
	"fmt"
	"time"
)

func handle(nServer int, itemCh chan int) chan Bid {
	bidCh := make(chan Bid)
	joinCh := make(chan int)
	for i := 0; i < nServer; i++ {
		go func() {
			for item := range itemCh {
				bidCh <- bid(item)
			}
			joinCh <- 1
		}()
	}

	go func() {
		for i := 0; i < nServer; i++ {
			<-joinCh
		}
		close(joinCh)
		close(bidCh)
	}()

	return bidCh
}

type Bid struct {
	item      int
	bidValue  int
	bidFailed bool
}

func generateInput(inCh chan int) {
	for i := 0; i < 15; i++ {
		inCh <- i
	}
	close(inCh)
}

func bid(item int) Bid {
	time.Sleep(time.Second * 5)
	return Bid{item, item + 1, false}
}

func main() {
	itemCh := make(chan int)

	go generateInput(itemCh)

	bidCh := handle(5, itemCh)
	for bid := range bidCh {
		fmt.Println("bid: ", bid)
	}
	fmt.Println("DONE!")
}