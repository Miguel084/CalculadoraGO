package main

import "testing"

func Test_divisao(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Divisão de 10 por 2", args{[]int{10, 2}}, 5},
		{"Divisão de 20 por 2 e 2", args{[]int{20, 2, 2}}, 5},
		{"Divisão por zero", args{[]int{10, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisao(tt.args.x...); got != tt.want {
				t.Errorf("divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplicacao(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Multiplicação de 2 e 3", args{[]int{2, 3}}, 6},
		{"Multiplicação de 2, 3 e 4", args{[]int{2, 3, 4}}, 24},
		{"Multiplicação sem argumentos", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplicacao(tt.args.x...); got != tt.want {
				t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Soma de 1, 2 e 5", args{[]int{1, 2, 5}}, 8},
		{"Soma de 10 e 5", args{[]int{10, 5}}, 15},
		{"Soma sem argumentos", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x...); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtracao(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Subtração de 10 e 5", args{[]int{10, 5}}, 5},
		{"Subtração de 20, 5 e 3", args{[]int{20, 5, 3}}, 12},
		{"Subtração sem argumentos", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtracao(tt.args.x...); got != tt.want {
				t.Errorf("subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}
