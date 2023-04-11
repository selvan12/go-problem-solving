package main

import "fmt"

// how garbage collector works in golang
// how channel know to establish the connection (send/recv) - high level architecture
// how memory stored in golang

// run gorotine to print the number from 1 to infinte times
// stop gorotine in 5 sec
const MaxSecondRun = 2 // 5 // in seconds

func simpleGoRoutine(chan1 chan string) {
	var val int
	for {
		select {
		case msg := <-chan1:
			fmt.Println("Stopping goroutine using chan ", msg)
			return
		default:
			fmt.Println(" ", val)
			val++
		}
	}
}

/*
func main() {

	fmt.Println("Test Program")

	// fmt.Println("stop go routine after certain duartion")
	// chan1 := make(chan string)
	// go simpleGoRoutine(chan1)

	// select {
	// case <-time.After(MaxSecondRun * time.Second):
	// 	fmt.Println("Sending quit from main!!!")
	// 	chan1 <- "quit"
	// 	fmt.Println("Stopping from main!!!")
	// 	return
	// }

	//fmt.Println("Out: ", sortedSquares([]int{-4, -2, 1, 3, 7, 9}))

}
*/
