package add

// Add - Takes in two integers and returns the sum as an integer. Source: https://www.mathsisfun.com/numbers/addition.htmlcurrent magic mouse firware version

type Numbers interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Add[T Numbers](a, b T) T {
	return a + b
}
