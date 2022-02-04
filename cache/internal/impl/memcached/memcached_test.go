package memcached

import (
	"context"
	"fmt"
	commontesting "github.com/cirruslabs/backbone-services/cache/internal/impl/testing"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

func TestMemcached(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:        "memcached:1.5.16-alpine",
		ExposedPorts: []string{"11211/tcp"},
		WaitingFor:   wait.ForListeningPort("11211/tcp"),
	}
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	ip, err := container.Host(context.Background())
	require.NoError(t, err)

	mappedPort, err := container.MappedPort(context.Background(), "11211")
	require.NoError(t, err)

	server, err := NewMemcached(
		fmt.Sprintf("%s:%s", ip, mappedPort.Port()),
	)
	require.NoError(t, err)

	t.Run("Memcached", func(t *testing.T) {
		commontesting.TestCacheImplementation(t, server)
	})
}
