package main

import "fmt"

func main() {

	// Задание 8
	var d uint16

	fmt.Scan(&d)
	fmt.Println("It is", d/30, "hours", d*12%360/6, "minutes")

}
