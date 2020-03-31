package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hazbo/calc-svc/pkg/server"
)

func main() {
	var port *string = flag.String(
		"p", "2020", "The port to listen on")

	flag.Parse()
	log.Fatal(server.New("tcp", fmt.Sprintf(":%s", *port)).Listen())
}
