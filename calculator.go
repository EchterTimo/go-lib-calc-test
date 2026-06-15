package calculator

// Number is a custom constraint that accepts both integers and floats.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Add takes two numbers of the same type and returns their sum.
func Add[T Number](a, b T) T {
	return a + b
}

// Subtract takes two numbers of the same type and returns their difference.
func Subtract[T Number](a, b T) T {
	return a - b
}

// Multiply takes two numbers of the same type and returns their product.
func Multiply[T Number](a, b T) T {
	return a * b
}

// Divide takes two numbers of the same type and returns their quotient.
func Divide[T Number](a, b T) T {
	return a / b
}
