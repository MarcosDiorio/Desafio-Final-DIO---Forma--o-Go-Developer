package main

import "fmt"

func main() {
	x := soma(88, 1, 741)
	y := multiplica(9, 150)
	v := divisao(90, 3)
	z := subtra(480, 22, 15)
	fmt.Println(x, y, v, z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divisao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total / v
	}
	return total
}
func subtra(i ...int) int {
	total := 0
	for _, v := range i {
		total -= v
	}
	return total
}
