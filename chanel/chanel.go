package main

import (
	"fmt"
	"time"
)

func write(ch chan<- int) {
	ch <- 1
}

func main() {
	ch := make(chan int)
	go write(ch)
	time.Sleep(1 * time.Second)
	fmt.Println(ch)
	close(ch)
}

//Способ паралельного ожидания сигналов из разных кналов:

func someFunc(chVal chan int, chErr chan error) {
	var i int
	chVal <- i

}

func twoChanListen() {
	chErr := make(chan error)
	chVal := make(chan int)

	go someFunc()

	switch {
	case <-chErr:
		panic(err)
	case <-chVal:
		fmt.Println(chVal)
	}
}
