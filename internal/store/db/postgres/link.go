package postgres

import (
	"context"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/go-pg/pg/v10"
)

type LinkRepository struct {
	db *pg.DB
}

// New Link Repository
func NewLinkRepository(db *pg.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (lr *LinkRepository) Get(ctx context.Context, hash string) (*entity.LinkDB, error) {
	link := &entity.LinkDB{}
	err := lr.db.Model(link).Where("hash = ?", hash).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return link, nil
}
func (lr *LinkRepository) New(ctx context.Context, link *entity.LinkDB) (*entity.LinkDB, error) {
	existingLink, err := lr.Get(ctx, link.Hash)
	if err != nil {
		return nil, err
	}
	if existingLink != nil {
		return existingLink, nil
	}
	_, err = lr.db.Model(link).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return link, err
}
