package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/stretchr/testify/assert"
)

//godotenv -f ./.env go test ./... -cover -v
func TestPostgresLinksIntegrational(t *testing.T) {

	db, err := NewPgDb(os.Getenv("PG_URL_TESTS"))
	if err != nil {
		panic("db connection failed")
	}
	repo := NewLinkRepository(db.DB)
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
			name:      "valid new",
			InputLink: &entity.LinkDB{Origin: "https://google.com", Hash: "Lhr4BWAi"},
			OutpuLink: &entity.LinkDB{Origin: "https://google.com", Hash: "Lhr4BWAi"},
			err:       nil,
		},
	}

	for _, tc := range testsNew {
		out, err := repo.New(ctx, tc.InputLink)
		if tc.err != nil {
			assert.Error(t, err)
		}
		if tc.OutpuLink != nil {
			assert.Equal(t, tc.InputLink.Hash, out.Hash)
			//assert.Equal(t, tc.OutpuLink, out)
		}
	}

	testsGet := []test{
		{
			name:      "valid get",
			hash:      "Lhr4BWAi",
			OutpuLink: &entity.LinkDB{Hash: "Lhr4BWAi"},
			err:       nil,
		},
		{
			name: "non existing hash",
			hash: "1234",
			err:  nil,
		},
	}

	for _, tg := range testsGet {
		out, err := repo.Get(ctx, tg.hash)
		if tg.err != nil {
			assert.Error(t, err)
		}
		if tg.OutpuLink != nil {
			assert.Equal(t, tg.hash, out.Hash)
		}
	}
}
