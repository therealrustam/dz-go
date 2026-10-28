package main

import (
	"fmt"
	"math/rand"
)

func producer(ch chan int) {
	numbers := make([]int, 10, 10)
	for index, _ := range numbers {
		numbers[index] = rand.Intn(100)
	}
	for _, number := range numbers {
		ch <- number
	}

}

func consumer(otherCh chan int, resultCh chan int) {
	for i := 0; i < 10; i++ {
		number := <-otherCh
		result := number * number
		resultCh <- result
	}
}

func main() {
	resultCh := make(chan int)
	otherCh := make(chan int)
	go producer(otherCh)
	go consumer(otherCh, resultCh)
	for i := 0; i < 10; i++ {
		result := <-resultCh
		fmt.Printf("%d ", result)
	}

}
