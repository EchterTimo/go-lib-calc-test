package main

import calc "github.com/EchterTimo/go-lib-calc-test"

func main() {
	value1 := 2
	value2 := 3
	result := calc.Add(value1, value2)
	println("The result of adding", value1, "and", value2, "is:", result)

	result2 := calc.Subtract(value1, value2)
	println("The result of subtracting", value2, "from", value1, "is:", result2)

	result3 := calc.Multiply(value1, value2)
	println("The result of multiplying", value1, "and", value2, "is:", result3)

	result4 := calc.Divide(value1, value2)
	println("The result of dividing", value1, "by", value2, "is:", result4)

}
