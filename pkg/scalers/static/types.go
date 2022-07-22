package static

import "sync"

type Scaler struct {
	down          bool
	readWriteLock sync.RWMutex
}

type scaleResponse struct {
	Name     string
	Replicas int64
	Region   string
	Primary  bool
}
