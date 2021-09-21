package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateNumber(numeros chan<- int) {
	defer close(numeros)
	for i :=0 ; i < 20; i++ {
		numeros <- rand.Int()
	}
}

func verifyNumber(numeros <-chan int, join *sync.WaitGroup) {
	defer join.Done()
	for element := range numeros {
		fmt.Println(element % 2 == 0)
	}
}

func main() {
	numeros := make(chan int)
	var join sync.WaitGroup
	join.Add(1)
	go generateNumber(numeros)
	go verifyNumber(numeros, &join)
	join.Wait()
}