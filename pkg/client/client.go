package client

import (
	"context"
	"fmt"

	"github.com/hazbo/calc-svc/pkg/calc"
	"google.golang.org/grpc"
)

// Client represents a client to be used with the Calculator Service containing
// the remote address which is then used to create the connection.
type Client struct {
	addr string
	conn *grpc.ClientConn
}

// Symbols for referring to one of the functions.
const (
	funcSum int = iota
	funcMul
	funcDiv

	numFuncs
)

// Symbols for storing the required number of args for each function.
const (
	funcSumArgc int = 2
	funcMulArgc int = 2
	funcDivArgc int = 2
)

// New returns a new instance of the client with a given remote address.
func New(addr string) *Client {
	return &Client{
		addr: addr,
		conn: nil,
	}
}

// Connect attempts to make a connection, returning an error if for whatever
// reason this cannot be done.
func (c *Client) Connect() error {
	var (
		conn *grpc.ClientConn
		err  error
	)

	conn, err = grpc.Dial(c.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("Error connecting to %s, %s\n", c.addr, err)
	}

	c.conn = conn

	return nil
}

// Sum provides a public interface to the internal callSum call.
func (c *Client) Sum(x, y int64) (int64, error) {
	var (
		z   interface{}
		err error
	)

	z, err = c.call(funcSum, funcSumArgc, x, y)

	return z.(int64), err
}

// Mul provides a public interface to the internal callMul call.
func (c *Client) Mul(x, y int64) (int64, error) {
	var (
		z   interface{}
		err error
	)

	z, err = c.call(funcMul, funcMulArgc, x, y)

	return z.(int64), err
}

// Div provides a public interface to the internal callDiv call.
func (c *Client) Div(x, y float64) (float64, error) {
	var (
		z   interface{}
		err error
	)

	z, err = c.call(funcDiv, funcDivArgc, x, y)

	return z.(float64), err
}

func (c *Client) Close() {
	c.conn.Close()
}

// call is for making a generic call in which code is executed before each
// actual call is made. The main thing here is first checking a connection has
// been previously made before attempting any call.
//
// The function ID passed as the first argument is one of the list of constants
// declared earlier on in this file.
func (c *Client) call(fn int, argc int, argv ...interface{}) (interface{}, error) {
	var (
		calcClient calc.CalculatorClient
	)

	// Always check to make sure the connection has first been made.
	if c.conn == nil {
		return nil, fmt.Errorf("Error: connection must first be made to remote")
	}

	// Check the correct amount of arguments have been passed.
	if len(argv) != argc {
		return nil, fmt.Errorf("Error: expecting %d args, got %d\n", argc, len(argv))
	}

	// Check the function ID actually exists.
	if fn >= numFuncs {
		return nil, fmt.Errorf("Error: function ID, '%d', is not valid", fn)
	}

	// TODO: move this out of here to make it testable
	calcClient = calc.NewCalculatorClient(c.conn)

	// Search for and then call the function
	switch fn {
	case funcSum:
		return c.callSum(calcClient, argv)
	case funcMul:
		return c.callMul(calcClient, argv)
	case funcDiv:
		return c.callDiv(calcClient, argv)
	}

	return nil, fmt.Errorf("Error: function ID, '%d', is not valid", fn)
}

// The internal calling functions
//
// In retrospect I think I should have started with functions that differ from
// each other, as these 3 all have similar func signatures, which is why the
// implementations are also similar.
//
// These functions can be used to perform individual checks on a given function
// call which unlike things done in `call`, will only effect a single function.

// callSum prepares the input to be remotely executed for the Sum function and
// then makes the call.
func (c *Client) callSum(
	calcClient calc.CalculatorClient, args []interface{}) (int64, error) {

	var (
		rep *calc.SumResponse
		err error
	)

	if rep, err = calcClient.Sum(context.Background(), &calc.SumRequest{
		X: args[0].(int64),
		Y: args[1].(int64),
	}); err != nil {
		return 0, err
	}

	return rep.GetZ(), nil
}

// callMul prepares the input to be remotely executed for the Mul function and
// then makes the call.
func (c *Client) callMul(
	calcClient calc.CalculatorClient, args []interface{}) (int64, error) {

	var (
		rep *calc.MulResponse
		err error
	)

	if rep, err = calcClient.Mul(context.Background(), &calc.MulRequest{
		X: args[0].(int64),
		Y: args[1].(int64),
	}); err != nil {
		return 0, err
	}

	return rep.GetZ(), nil
}

// callDiv prepares the input to be remotely executed for the Div function and
// then makes the call.
func (c *Client) callDiv(
	calcClient calc.CalculatorClient, args []interface{}) (float64, error) {

	var (
		rep *calc.DivResponse
		err error
	)

	if rep, err = calcClient.Div(context.Background(), &calc.DivRequest{
		X: args[0].(float64),
		Y: args[1].(float64),
	}); err != nil {
		return 0, err
	}

	return rep.GetZ(), nil
}
