package main

import (
	"fmt"
	"math"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func max(args ...int) int {
	var maxNum int = math.MinInt32

	for _, num := range args {
		if num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}

func generadorImpares() func() uint {
	i := uint(1)

	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}

func intercambia(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func main() {
	fmt.Println("fibonacci")
	for i := 0; i < 11; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println()

	fmt.Println("max")
	fmt.Println(max(23, 243, 123, 654, 23))
	fmt.Println()

	fmt.Println("generator A invocations")
	var myGenerator = generadorImpares()
	fmt.Println(myGenerator())
	fmt.Println(myGenerator())
	fmt.Println(myGenerator())
	fmt.Println("generator B invocations")
	var myOtherGenerator = generadorImpares()
	fmt.Println(myOtherGenerator())
	fmt.Println(myOtherGenerator())
	fmt.Println(myOtherGenerator())
	fmt.Println(myOtherGenerator())
	fmt.Println("generator A invocation")
	fmt.Println(myGenerator())
	fmt.Println()

	fmt.Println("intercambio con punteros")
	a := int(0)
	b := int(1)
	fmt.Printf("antes: a: %d, b: %d\n", a, b)
	intercambia(&a, &b)
	fmt.Printf("despues: a: %d, b: %d", a, b)
}
