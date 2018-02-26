package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func getNum() int {
	num := 1

	fmt.Print("Enter the total Fibonnacci numbers to print: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println(err)
	}

	return num
}

func main() {
	f := fib()
	num := getNum()

	for i := 0; i < num; i++ {
		fmt.Println(f())
	}
}
