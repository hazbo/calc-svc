package mock_calc_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	calc "github.com/hazbo/calc-svc/pkg/calc"
	mcalc "github.com/hazbo/calc-svc/pkg/mock_calc"
)

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculatorClient := mcalc.NewMockCalculatorClient(ctrl)

	// TODO: switch out these gomock.Any for the matcher interface ^
	mockCalculatorClient.EXPECT().
		Sum(gomock.Any(),
			gomock.Any()).Return(&calc.SumResponse{Z: 30}, nil)
	testSum(t, mockCalculatorClient)
}

func testSum(t *testing.T, client calc.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var expecting int64 = 30

	r, err := client.Sum(ctx, &calc.SumRequest{X: 10, Y: 20})
	if err != nil || r.Z != expecting {
		t.Errorf("Error: %s\nExpecting %d, got %d\n", err, expecting, r.Z)
	}
}

func TestMul(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculatorClient := mcalc.NewMockCalculatorClient(ctrl)

	// TODO: switch out these gomock.Any for the matcher interface ^
	mockCalculatorClient.EXPECT().
		Mul(gomock.Any(),
			gomock.Any()).Return(&calc.MulResponse{Z: 30}, nil)
	testMul(t, mockCalculatorClient)
}

func testMul(t *testing.T, client calc.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var expecting int64 = 30

	r, err := client.Mul(ctx, &calc.MulRequest{X: 3, Y: 10})
	if err != nil || r.Z != expecting {
		t.Errorf("Error: %s\nExpecting %d, got %d\n", err, expecting, r.Z)
	}
}

func TestDiv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculatorClient := mcalc.NewMockCalculatorClient(ctrl)

	// TODO: switch out these gomock.Any for the matcher interface ^
	mockCalculatorClient.EXPECT().
		Div(gomock.Any(),
			gomock.Any()).Return(&calc.DivResponse{Z: 30}, nil)
	testDiv(t, mockCalculatorClient)
}

func testDiv(t *testing.T, client calc.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var expecting float64 = 30

	r, err := client.Div(ctx, &calc.DivRequest{X: 90, Y: 3})
	if err != nil || r.Z != expecting {
		t.Errorf("Error: %s\nExpecting %f, got %f\n", err, expecting, r.Z)
	}
}
