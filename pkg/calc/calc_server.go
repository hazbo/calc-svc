package calc

import (
	"context"
)

// calculatorServer implements CalculatorServer and contains only a Calculator.
type calculatorServer struct {
	c Calculator
}

// NewServer returns a new calculator server with a given Calculator.
func NewServer(c Calculator) CalculatorServer {
	return &calculatorServer{c: c}
}

// NewDefaultServer returns a new calculator server with the default builtin
// calculator.
func NewDefaultServer() CalculatorServer {
	return NewServer(New())
}

// Sum makes a call to the calculator's internal Sum function and returns the
// result within a SumResponse object.
func (cs *calculatorServer) Sum(
	ctx context.Context,
	req *SumRequest) (*SumResponse, error) {

	return &SumResponse{Z: cs.c.Sum(req.GetX(), req.GetY())}, nil
}

// Mul makes a call to the calculator's internal Mul function and returns the
// result within a MulResponse object.
func (cs *calculatorServer) Mul(
	ctx context.Context,
	req *MulRequest) (*MulResponse, error) {

	return &MulResponse{Z: cs.c.Mul(req.GetX(), req.GetY())}, nil
}

// Div makes a call to the calculator's internal Div function and returns the
// result within a DivResponse object.
func (cs *calculatorServer) Div(
	ctx context.Context,
	req *DivRequest) (*DivResponse, error) {

	return &DivResponse{Z: cs.c.Div(req.GetX(), req.GetY())}, nil
}
