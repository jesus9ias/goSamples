package main

import (
	"fmt"
)

func main() {
	c := 45
	fmt.Println(c, &c)
	c = 50
	fmt.Println(c, &c)
	d := &c
	c = 60
	fmt.Println(*d, d)
}
