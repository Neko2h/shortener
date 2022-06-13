package memory

import (
	"testing"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/stretchr/testify/assert"
)

/*
go test -coverprofile coverage.out ./...
go tool cover -html coverage.out
*/

func TestMemoryIntegrational(t *testing.T) {
	testLink := &entity.LinkDB{
		Origin: "https://google.com", Hash: "Lhr4BWAi",
	}
	added := addElementToStorage(testLink)
	found := sliceSearch(InmemoryStorage, "Lhr4BWAi")
	assert.Equal(t, added.Hash, found.Hash)
	assert.Equal(t, added.Origin, found.Origin)
	notFound := sliceSearch(InmemoryStorage, "")
	assert.Nil(t, notFound)
}
