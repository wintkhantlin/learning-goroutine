package main

import (
	"fmt"
	"math/big"
)

func fib(n int, c chan *big.Int) {
	if n <= 1 {
		c <- big.NewInt(int64(n))
		return
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	tmp := new(big.Int)
	for i := 2; i <= n; i++ {
		tmp.Set(b)
		b.Add(a, b)
		a.Set(tmp)
	}
	c <- b
}

func main() {
	c := make(chan *big.Int)

	go fib(1444, c)
	go fib(1444, c)
	go fib(1444, c)
	go fib(1444, c)
	go fib(1444, c)
	go fib(1444, c)

	x, y, z, a, b, d := <-c, <-c, <-c, <-c, <-c, <-c

	fmt.Println(x)
	fmt.Println("----")
	fmt.Println(y)
	fmt.Println("----")
	fmt.Println(z)
	fmt.Println("----")
	fmt.Println(a)
	fmt.Println("----")
	fmt.Println(b)
	fmt.Println("----")
	fmt.Println(d)
}
