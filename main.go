package main

import "fmt"

func main() {
	x := soma(1, 2, 5)
	y := subtracao(10, 5)
	z := multiplicacao(2, 3)
	w := divisao(10, 2)
	fmt.Println(x, y, z, w)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func subtracao(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	subtracao := x[0]
	for _, v := range x[1:] {
		subtracao -= v
	}
	return subtracao
}

func multiplicacao(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	multiplicacao := 1
	for _, v := range x {
		multiplicacao *= v
	}
	return multiplicacao
}

func divisao(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	divisao := x[0]
	for _, v := range x[1:] {
		if v == 0 {
			return 0
		}
		divisao /= v
	}
	return divisao
}
