package main

import (
	"fmt"
	"math"
)

func main() {

	// Задание 11
	const pi float64 = 3.141592
	var a, b, s, v float64

	fmt.Scan(&a, &b)

	s = pi * math.Sqrt(a) * math.Sqrt(b)
	v = (4.0 / 3) * pi * math.Sqrt(a) * b

	fmt.Println("Площадь поверхности:", s, "\nОбъем тела:", v)

}
