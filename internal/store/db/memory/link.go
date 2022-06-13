package memory

import (
	"context"

	"github.com/Neko2h/shortener/internal/entity"
)

type LinkRepository struct {
	*Memory
}

//New Link Repository
func NewLinkRepository(db *Memory) *LinkRepository {
	return &LinkRepository{db}
}

func (mdb *Memory) Get(ctx context.Context, hash string) (*entity.LinkDB, error) {

	mdb.Lock()
	link := sliceSearch(InmemoryStorage, hash)
	mdb.Unlock()

	return link, nil
}

func (mdb *Memory) New(ctx context.Context, link *entity.LinkDB) (*entity.LinkDB, error) {

	existingLink, _ := mdb.Get(ctx, link.Hash)

	if existingLink != nil {
		return existingLink, nil
	}

	newLink := addElementToStorage(link)
	mdb.Unlock()

	return newLink, nil
}
