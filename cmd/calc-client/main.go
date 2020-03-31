package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hazbo/calc-svc/pkg/client"
)

func main() {
	var (
		r    *string
		c    *client.Client
		args []string
	)

	r = flag.String(
		"r", "127.0.0.1:2020", "The IP and port of the remote")

	flag.Parse()
	args = flag.Args()

	if len(args) < 3 {
		errExit("Expecting [func] Operand1 Operand2")
	}

	c = client.New(*r)

	if err := c.Connect(); err != nil {
		errExit("Unable to make connection to remote")
	}

	defer c.Close()

	switch args[0] {
	case "sum":
		callSum(c, args)
		break
	case "mul":
		callMul(c, args)
		break
	case "div":
		callDiv(c, args)
		break
	default:
		fmt.Printf("Call must be one of 'sum', 'mul', 'div'\n")
	}
}

func callSum(c *client.Client, args []string) {
	var (
		x, y int
		err  error
		z    int64
	)

	if x, err = strconv.Atoi(args[1]); err != nil {
		errExit("Invalid first operand")
	}

	if y, err = strconv.Atoi(args[2]); err != nil {
		errExit("Invalid second operand")
	}

	if z, err = c.Sum(int64(x), int64(y)); err != nil {
		errExit("Unable to perform remote call")
	}

	fmt.Printf("%d\n", z)
}

func callMul(c *client.Client, args []string) {
	var (
		x, y int
		err  error
		z    int64
	)

	if x, err = strconv.Atoi(args[1]); err != nil {
		errExit("Invalid first operand")
	}

	if y, err = strconv.Atoi(args[2]); err != nil {
		errExit("Invalid second operand")
	}

	if z, err = c.Mul(int64(x), int64(y)); err != nil {
		errExit("Unable to perform remote call")
	}

	fmt.Printf("%d\n", z)
}

func callDiv(c *client.Client, args []string) {
	var (
		x, y float64
		err  error
		z    float64
	)

	if x, err = strconv.ParseFloat(args[1], 64); err != nil {
		errExit("Invalid first operand")
	}

	if y, err = strconv.ParseFloat(args[2], 64); err != nil {
		errExit("Invalid second operand")
	}

	if z, err = c.Div(x, y); err != nil {
		errExit("Unable to perform remote call")
	}

	fmt.Printf("%f\n", z)
}

func errExit(msg string) {
	fmt.Printf("Error: %s\n", msg)
	os.Exit(1)
}
