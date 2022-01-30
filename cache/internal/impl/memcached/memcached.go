package memcached

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/cirruslabs/backbone-services/proto/cache"
	"google.golang.org/protobuf/types/known/emptypb"
)

type memcached struct {
	client *memcache.Client
}

func NewMemcached(server ...string) (cache.CacheServer, error) {
	client := memcache.New(server...)
	err := client.Ping()
	if err != nil {
		return nil, err
	}
	return &memcached{
		client: client,
	}, nil
}

func (m memcached) Get(ctx context.Context, request *cache.GetRequest) (*cache.GetResponse, error) {
	items, err := m.client.GetMulti(request.Keys)
	if err != nil {
		return nil, err
	}
	entries := make(map[string][]byte, len(items))
	for key, item := range items {
		entries[key] = item.Value
	}
	return &cache.GetResponse{
		CacheEntries: entries,
	}, nil
}

func (m memcached) Put(ctx context.Context, request *cache.PutRequest) (*emptypb.Empty, error) {
	var lastError error
	for key, bytes := range request.GetEntries() {
		err := m.client.Set(&memcache.Item{
			Key:        key,
			Value:      bytes,
			Expiration: request.ExpiresInSeconds,
		})
		if err != nil {
			lastError = err
		}
	}
	if lastError != nil {
		return nil, lastError
	}
	return &emptypb.Empty{}, nil
}

func (m memcached) CompareAndSet(ctx context.Context, request *cache.CASRequest) (*cache.CASResponse, error) {
	panic("implement me")
}

func (m memcached) CompareAndDelete(ctx context.Context, request *cache.CADRequest) (*cache.CADResponse, error) {
	panic("implement me")
}

func (m memcached) MultiDelete(ctx context.Context, request *cache.MultiDeleteRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (m memcached) RateLimit(ctx context.Context, request *cache.RateLimitRequest) (*cache.RateLimitResponse, error) {
	panic("implement me")
}

func (m memcached) Inc(ctx context.Context, request *cache.IncRequest) (*cache.CounterResponse, error) {
	panic("implement me")
}

func (m memcached) Dec(ctx context.Context, request *cache.DecRequest) (*cache.CounterResponse, error) {
	panic("implement me")
}
