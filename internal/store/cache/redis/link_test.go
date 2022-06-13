package redis

import (
	"context"
	"os"
	"testing"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/stretchr/testify/assert"
)

/*
go test -coverprofile coverage.out ./...
go tool cover -html coverage.out

*/
func TestRedisLinksIntegrational(t *testing.T) {
	t.Setenv("REDIS_HOST", "localhost:6379")
	cache, _ := NewRedisClient(os.Getenv("REDIS_HOST"))
	repo := NewLinkCache(cache.Client)
	ctx := context.Background()

	type test struct {
		name      string
		InputLink *entity.LinkDB
		OutpuLink *entity.LinkDB
		err       error
		hash      string
	}

	testsNew := []test{
		{
			name:      "valid set",
			InputLink: &entity.LinkDB{Origin: "https://google.com", Hash: "Lhr4BWAi"},
			OutpuLink: &entity.LinkDB{Origin: "https://google.com", Hash: "Lhr4BWAi"},
			err:       nil,
		},
	}

	testsGet := []test{
		{
			name: "valid get",
			hash: "Lhr4BWAi",
			err:  nil,
		},
		{
			name: "non existing get",
			hash: "",
			err:  nil,
		},
	}

	for _, tc := range testsNew {
		t.Logf("running: %s", tc.name)
		err := repo.Set(ctx, tc.InputLink)
		if tc.err != nil {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

	for _, tc := range testsGet {
		t.Logf("running: %s", tc.name)
		_, err := repo.Get(ctx, tc.hash)
		if tc.err != nil {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

}
