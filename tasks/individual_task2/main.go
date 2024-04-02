package main

import (
	"fmt"
	"math"
)

func main() {

	//Индивидуальное задание 2. Вариант 3
	var a, b, c, p, s float64

	fmt.Print("Введите a, b и c:")
	fmt.Scan(&a, &b, &c)

	p = (a + b + c) / 2

	s = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Println("Площадь треугольника равна:", s)

}
