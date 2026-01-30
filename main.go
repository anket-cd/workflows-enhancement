package main

import (
	"anket-cd/we/mathsoprn/multi"
	"anket-cd/we/mathsoprn/substraction"
	"anket-cd/we/mathsoprn/sum"
	"fmt"
)

func main() {
	a, b := 10, 5
	fmt.Println(a, "+", b, "Sum is : ", sum.Sum(a, b))
	//---
	fmt.Println(a, "-", b, "Subtraction is :", substraction.Substraction(a, b))
	//---
	fmt.Println(a, "*", b, "Multiplication is :", multi.MultiPlication(a, b))
}
