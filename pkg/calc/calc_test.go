package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	c := New()

	var (
		exr1 int64 = 10
		exr2 int64 = 50
		exr3 int64 = 1003
	)

	r1 := c.Sum(3, 7)
	if r1 != exr1 {
		t.Errorf("Expecting %d, got %d\n", exr1, r1)
	}

	r2 := c.Sum(30, 20)
	if r2 != exr2 {
		t.Errorf("Expecting %d, got %d\n", exr2, r2)
	}

	r3 := c.Sum(999, 4)
	if r3 != exr3 {
		t.Errorf("Expecting %d, got %d\n", exr3, r3)
	}
}

func TestMul(t *testing.T) {
	c := New()

	var (
		exr1 int64 = 10
		exr2 int64 = 50
		exr3 int64 = 1003
	)

	r1 := c.Mul(2, 5)
	if r1 != exr1 {
		t.Errorf("Expecting %d, got %d\n", exr1, r1)
	}

	r2 := c.Mul(2, 25)
	if r2 != exr2 {
		t.Errorf("Expecting %d, got %d\n", exr2, r2)
	}

	r3 := c.Mul(1, 1003)
	if r3 != exr3 {
		t.Errorf("Expecting %d, got %d\n", exr3, r3)
	}
}

func TestDiv(t *testing.T) {
	c := New()

	var (
		exr1 float64 = 10
		exr2 float64 = 50
		exr3 float64 = 1003
	)

	r1 := c.Div(30, 3)
	if r1 != exr1 {
		t.Errorf("Expecting %f, got %f\n", exr1, r1)
	}

	r2 := c.Div(500, 10)
	if r2 != exr2 {
		t.Errorf("Expecting %f, got %f\n", exr2, r2)
	}

	r3 := c.Div(2006, 2)
	if r3 != exr3 {
		t.Errorf("Expecting %f, got %f\n", exr3, r3)
	}
}

func TestTest(t *testing.T) {
	c := New()

	if !c.Test() {
		t.Error("Calculator internal test should always be true")
	}
}
