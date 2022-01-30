package testing

import (
	"context"
	"github.com/cirruslabs/backbone-services/proto/cache"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestCacheImplementation(t *testing.T, srv cache.CacheServer) {
	l := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	cache.RegisterCacheServer(grpcServer, srv)
	go func() {
		_ = grpcServer.Serve(l)
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
			return l.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := cache.NewCacheClient(conn)

	t.Run("Get", func(t *testing.T) {
		testGet(t, client)
	})
}

func testGet(t *testing.T, client cache.CacheClient) {
	_, err := client.Put(context.Background(), &cache.PutRequest{
		Entries: map[string][]byte{
			"foo": []byte("bar"),
		},
		ExpiresInSeconds: 60,
	})
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Get(context.Background(), &cache.GetRequest{
		Keys: []string{"foo"},
	})
	if err != nil {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []byte("bar"), resp.GetCacheEntries()["foo"])
}
