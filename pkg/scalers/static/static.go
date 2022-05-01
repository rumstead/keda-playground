package static

import (
	"context"
	"log"
	"math"
	"sync"
	"time"
)
import pb "keda-cnp-scaler/pkg/scalers/protos"

type Scaler struct {
	down          bool
	readWriteLock sync.RWMutex
}

func (s *Scaler) Swap() bool {
	s.readWriteLock.Lock()
	defer s.readWriteLock.Unlock()
	s.down = !s.down
	log.Printf("swap: %t\n", s.down)
	return !s.down
}

func (s *Scaler) Down() bool {
	s.readWriteLock.RLock()
	defer s.readWriteLock.RUnlock()
	return s.down
}

func (s *Scaler) IsActive(_ context.Context, _ *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	log.Printf("IsActive: %t\n", s.Down())
	return &pb.IsActiveResponse{
		Result: s.Down(),
	}, nil
}

func (s *Scaler) StreamIsActive(scaledObj *pb.ScaledObjectRef, epsServer pb.ExternalScaler_StreamIsActiveServer) error {
	for {
		select {
		case <-epsServer.Context().Done():
			log.Println("Call cancelled...")
			// call cancelled
			return nil
		case <-time.Tick(5 * time.Millisecond):
			active, err := s.IsActive(context.TODO(), scaledObj)
			if err != nil {
				return err
			}
			err = epsServer.Send(&pb.IsActiveResponse{
				Result: active.Result,
			})
			return err
		}
	}
}

func (s *Scaler) GetMetricSpec(context.Context, *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	return &pb.GetMetricSpecResponse{
		MetricSpecs: []*pb.MetricSpec{{
			MetricName: "",
			TargetSize: int64(math.MaxInt),
		}},
	}, nil
}

func (s *Scaler) GetMetrics(_ context.Context, _ *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return nil, nil
}

func NewStaticScaler() *Scaler {
	return &Scaler{
		down: false,
	}
}
