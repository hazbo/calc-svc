package calc

// Calculator represents the behaviour of a very simple calculator with only 4
// functions as of now, including a testing function.
type Calculator interface {
	Sum(x, y int64) int64
	Mul(x, y int64) int64
	Div(x, y float64) float64
	Test() bool
}

// calculator represents the internal builtin calculator that is used for the
// server by default by implementing Calculator.
type calculator struct {
}

// New returns a pointer to a new calculator.
func New() *calculator {
	return &calculator{}
}

// Sum takes input x and y and then returns the sum of the values.
func (c *calculator) Sum(x, y int64) int64 {
	return x + y
}

// Mul takes input x and y and then returns the product of the values.
func (c *calculator) Mul(x, y int64) int64 {
	return x * y
}

// Div takes input x and y and then returns the division of the values.
func (c *calculator) Div(x, y float64) float64 {
	return x / y
}

// Test performs a number of basic tests to ensure that the calculator is
// perfoming correctly. You can think of these as sort-of runtime unit tests.
// In this calculator we're just using Go arithmetic operators so we can be
// fairly sure that these tests will pass, however if there were to be more
// complex functions or functions that invoke side-effects, running some basic
// tests at runtime in conjunction with the grpc health checks, means we can be
// fairly confident that the calculator is performing correctly at runtime.
//
// If you're implementing your own Calculator, you may choose to write your own
// tests in the Test func, or to disable this just return true (PASS).
func (c *calculator) Test() bool {
	if c.Sum(20, 30) != 50 {
		return false
	}

	return true
}
