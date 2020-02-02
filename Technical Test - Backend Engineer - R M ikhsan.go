package main

import (
	"fmt"
	"strconv"
)

func Sums(x int, y int) int {
	return x + y
}
func Multiples(x int, y int) int {
	return x * y
}
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Prime (n int){
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		print(prime, " , ")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func FibonacciGen(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciGen(n-1) + FibonacciGen(n-2)
}
func Finobacci(n int){
	for i := 0; i < n; i++ {
		fmt.Print(strconv.Itoa(FibonacciGen(i)) + " ")
	}
}

func main() {
	fmt.Println(Sums(1, 2))
	fmt.Println("")
	fmt.Println(Multiples(1, 2))
	fmt.Println("")
	Prime(4)
	fmt.Println("")
	Finobacci(4)

}
