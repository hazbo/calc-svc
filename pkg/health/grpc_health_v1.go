package grpc_health_v1

import (
	"context"

	"github.com/hazbo/calc-svc/pkg/calc"
)

// healthServer implements HealthServer
type healthServer struct {
}

// NewServer returns the detault implementation of HealthServer.
func NewServer() HealthServer {
	return &healthServer{}
}

// Check performs the calculators internal tests to ensure that it is
// functioning correctly.
func (s *healthServer) Check(
	ctx context.Context,
	req *HealthCheckRequest) (*HealthCheckResponse, error) {

	// Perform a quick test on the calculator to make sure the results it
	// returns are accurate.
	switch calc.New().Test() {
	case true:
		return &HealthCheckResponse{
			Status: HealthCheckResponse_SERVING,
		}, nil
	default:
		return &HealthCheckResponse{
			Status: HealthCheckResponse_UNKNOWN,
		}, nil
	}
}

// Watch streams a SERVING result when probed.
func (s *healthServer) Watch(
	req *HealthCheckRequest,
	stream Health_WatchServer) error {

	stream.Send(&HealthCheckResponse{
		Status: HealthCheckResponse_SERVING,
	})

	return nil
}
