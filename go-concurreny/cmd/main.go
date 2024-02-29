package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var state int32
	for i := 0; i < 6; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Print(state)
}

// func printFunc(index int, wg *sync.WaitGroup, numberch chan int) {
// 	time.Sleep(1 * time.Second)
// 	numberch <- index
// 	wg.Done()
// }

// func check() {
// 	numberch := make(chan int, 128)

// 	wg := &sync.WaitGroup{}
// 	for i := range 100 {
// 		go printFunc(i, wg, numberch)
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// 	close(numberch)
// 	for ele := range numberch {
// 		fmt.Println(ele)
// 	}
// }
