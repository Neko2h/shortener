package redis

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//localhost:6379

func TestRedisIntegrational(t *testing.T) {
	t.Setenv("REDIS_HOST", "localhost:999999")
	_, err := NewRedisClient(os.Getenv("REDIS_HOST"))
	assert.Error(t, err)
}
