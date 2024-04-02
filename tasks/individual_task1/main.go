package main

import "fmt"

func main() {

	// Индивидуальное задание 1. Вариант 3
	var length, width, height, volume float64

	fmt.Print("Введите длину, ширину основания и высоту пирамиды: ")
	fmt.Scan(&length, &width, &height)

	volume = (length * width * height) / 3

	fmt.Print("Объём пирамиды равен:", volume)

}
