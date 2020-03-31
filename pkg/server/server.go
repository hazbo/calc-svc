package server

import (
	"fmt"
	"net"

	"github.com/hazbo/calc-svc/pkg/calc"
	"github.com/hazbo/calc-svc/pkg/health"
	"google.golang.org/grpc"
)

// server represents an abstract server containing the main grpc server along
// with the configuration.
type server struct {
	svr      *grpc.Server
	protocol string
	port     string
}

// New returns a new pointer to a server containing the given protocol and
// port.
func New(protocol, port string) *server {
	return &server{protocol: protocol, port: port}
}

// Listen attempts to listen on the configured port and then registers the
// needed servers, in this case this just being the actual calculator server
// and the healthcheck server.
func (s *server) Listen() error {
	var (
		ln  net.Listener
		err error
	)

	ln, err = net.Listen(s.protocol, s.port)
	if err != nil {
		return fmt.Errorf("Error: Unable to start server: %s\n", err)
	}

	defer ln.Close()

	s.svr = grpc.NewServer()
	s.registerServers()

	fmt.Printf("Server listening on port %s\n", s.port)

	return s.svr.Serve(ln)
}

// registerServers registers each individual server that is used.
func (s *server) registerServers() {
	// Register a new Calculator server
	calc.RegisterCalculatorServer(s.svr, calc.NewDefaultServer())

	// Register the health check server
	grpc_health_v1.RegisterHealthServer(s.svr, grpc_health_v1.NewServer())
}
