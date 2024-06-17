package main

import (
	"fmt"

	io "learngo.com/inputOutput"
)

var number = 12

func init() {
	fmt.Println("Initializing init() one")
	prev := number
	number = 5
	fmt.Println("Previous number var: ", prev, " has changed it value to: ", number)
}

func init() {
	fmt.Println("Initializing init() two")
}

func main() {
	// go print("main")
	// go print("other")
	// go print("testing concurrency")
	// go print("go is cool")
	// time.Sleep(time.Second)

	io.PrintData(1, io.Url)
	fmt.Println(io.Math(2, 5.2))
	io.Slices()
}
