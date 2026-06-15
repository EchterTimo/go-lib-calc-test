package main

import calc "github.com/EchterTimo/go-lib-calc-test"

func main() {
	value1 := 2
	value2 := 3
	result := calc.Add(value1, value2)
	println("The result of adding", value1, "and", value2, "is:", result)
}
