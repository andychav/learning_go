package andy_math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes in two ints and adds them together
// [https://www.mathsisfun.com/numbers/addition.html]:Addition Explanation
func Add[T Number](x, y T) T {
	return x + y
}
