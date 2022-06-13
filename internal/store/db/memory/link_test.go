package memory

import (
	"context"
	"testing"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestMemoryLinkIntegrational(t *testing.T) {

	testLink := &entity.LinkDB{
		Origin: "https://google.com", Hash: "Lhr4BWAi",
	}

	// go test -coverprofile coverage.out ./...
	// go tool cover -html coverage.out
	store := NewInMemoryDb()
	repo := NewLinkRepository(store)
	ctx := context.Background()

	type test struct {
		name      string
		inputLink *entity.LinkDB
		outpuLink *entity.LinkDB
		err       error
	}

	testsNew := []test{
		{
			name:      "new non existing",
			inputLink: testLink,
			outpuLink: testLink,
			err:       nil,
		},
	}

	for _, tg := range testsNew {
		t.Logf("running: %s", tg.name)
		out, _ := repo.New(ctx, tg.inputLink)

		if tg.outpuLink != nil {
			assert.Equal(t, tg.inputLink.Hash, out.Hash)
		}
	}
}
