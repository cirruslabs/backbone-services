package impl

import (
	"context"
	"github.com/cirruslabs/backbone-services/proto/cache"
	"google.golang.org/protobuf/types/known/emptypb"
)

type inmemory struct {
}

func NewInmemory() *inmemory {
	return &inmemory{}
}

func (i inmemory) Get(ctx context.Context, request *cache.GetRequest) (*cache.GetResponse, error) {
	panic("implement me")
}

func (i inmemory) Put(ctx context.Context, request *cache.PutRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (i inmemory) CompareAndSet(ctx context.Context, request *cache.CASRequest) (*cache.CASResponse, error) {
	panic("implement me")
}

func (i inmemory) CompareAndDelete(ctx context.Context, request *cache.CADRequest) (*cache.CADResponse, error) {
	panic("implement me")
}

func (i inmemory) MultiDelete(ctx context.Context, request *cache.MultiDeleteRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (i inmemory) RateLimit(ctx context.Context, request *cache.RateLimitRequest) (*cache.RateLimitResponse, error) {
	panic("implement me")
}

func (i inmemory) Inc(ctx context.Context, request *cache.IncRequest) (*cache.CounterResponse, error) {
	panic("implement me")
}

func (i inmemory) Dec(ctx context.Context, request *cache.DecRequest) (*cache.CounterResponse, error) {
	panic("implement me")
}
