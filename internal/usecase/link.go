package usecase

import (
	"context"
	"errors"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/Neko2h/shortener/internal/helpers"
	"github.com/Neko2h/shortener/internal/store/cache"
	"github.com/Neko2h/shortener/internal/store/db"
)

type LinkWebUsecase struct {
	db          db.LinkRepository
	cacheDriver cache.LinkCache
}

func NewLinkUsecase(ur db.LinkRepository, cacheDriver cache.LinkCache) *LinkWebUsecase {
	return &LinkWebUsecase{
		cacheDriver: cacheDriver,
		db:          ur,
	}
}

func (us *LinkWebUsecase) Get(ctx context.Context, hash string) (*entity.Link, error) {

	cachedLink, err := us.cacheDriver.Get(ctx, hash)
	if err != nil {
		return nil, err
	}
	if cachedLink == nil {
		link, err := us.db.Get(ctx, hash)
		if err != nil {
			return nil, err
		}
		if link == nil {
			return nil, nil
		}

		err = us.cacheDriver.Set(ctx, link)
		if err != nil {
			return nil, err
		}

		return link.ToWeb(), nil
	}

	return cachedLink.ToWeb(), nil
}

func (us *LinkWebUsecase) New(ctx context.Context, link *entity.Link) (*entity.Link, error) {

	if link.Origin == "" {
		return nil, errors.New("origin is empty")
	}

	hash := helpers.GenerateShortLink(link.Origin)

	link.Hash = hash

	newLink, err := us.db.New(ctx, link.ToDB())
	if err != nil {
		return nil, err
	}

	err = us.cacheDriver.Set(ctx, newLink)
	if err != nil {
		return nil, err
	}

	return newLink.ToWeb(), nil
}
