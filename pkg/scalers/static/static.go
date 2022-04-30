package static

import "context"
import pb "keda-cnp-scaler/pkg/scalers/protos"

type Scaler struct {
}

func (e *Scaler) IsActive(ctx context.Context, scaledObject *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	return nil, nil
}

func (e *Scaler) StreamIsActive(*pb.ScaledObjectRef, pb.ExternalScaler_StreamIsActiveServer) error {
	return nil
}

func (e *Scaler) GetMetricSpec(context.Context, *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	return nil, nil
}

func (e *Scaler) GetMetrics(context.Context, *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return nil, nil
}
