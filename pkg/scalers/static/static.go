package static

import (
	"context"
	"sync"
	"time"
)
import pb "keda-cnp-scaler/pkg/scalers/protos"

type Scaler struct {
	down          bool
	readWriteLock sync.RWMutex
}

func (s *Scaler) Swap() {
	s.readWriteLock.Lock()
	defer s.readWriteLock.Unlock()
	s.down = !s.down
}

func (s *Scaler) Down() bool {
	s.readWriteLock.RLock()
	defer s.readWriteLock.Unlock()
	return s.down
}

func (s *Scaler) IsActive(_ context.Context, _ *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	return &pb.IsActiveResponse{
		Result: s.Down(),
	}, nil
}

func (s *Scaler) StreamIsActive(_ *pb.ScaledObjectRef, epsServer pb.ExternalScaler_StreamIsActiveServer) error {
	for {
		select {
		case <-epsServer.Context().Done():
			// call cancelled
			return nil
		case <-time.Tick(time.Hour * 1):
			err := epsServer.Send(&pb.IsActiveResponse{
				Result: s.Down(),
			})
			return err
		}
	}
}

func (s *Scaler) GetMetricSpec(context.Context, *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	return &pb.GetMetricSpecResponse{}, nil
}

func (s *Scaler) GetMetrics(context.Context, *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return &pb.GetMetricsResponse{}, nil
}

func NewStaticScaler() *Scaler {
	return &Scaler{
		down: false,
	}
}
