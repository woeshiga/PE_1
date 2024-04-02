package main

import "fmt"

func main() {

	// Задание 4
	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)

	a = a * a
	b = b * b
	c = a + b

	fmt.Println(c)

}
