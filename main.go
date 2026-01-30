package main

import (
	"anket-cd/we/multiplication"
	"anket-cd/we/substraction"
	"anket-cd/we/sum"
	"fmt"
)

func main() {
	a, b := 10, 5
	fmt.Println(a, "+", b, "Sum is : ", sum.Sum(a, b))
	//---
	fmt.Println(a, "-", b, "Subtraction is :", substraction.Substraction(a, b))
	//---
	fmt.Println(a, "*", b, "Multiplication is :", multiplication.MultiPlication(a, b))
}
