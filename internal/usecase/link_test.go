package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Neko2h/shortener/internal/entity"
	mockCache "github.com/Neko2h/shortener/internal/store/cache/mocks"
	mockDb "github.com/Neko2h/shortener/internal/store/db/mocks"
	"github.com/stretchr/testify/assert"
)

/*
go test -coverprofile coverage.out ./...
go tool cover -html coverage.out
*/

func TestLinkUnitUsecase(t *testing.T) {

	input := &entity.Link{
		Origin: "https://google.com",
		Hash:   "Lhr4BWAi",
	}

	tests := []struct {
		name          string
		expectactions func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache)
		input         *entity.Link
		callType      string
		err           error
	}{
		{
			name: "new and valid",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkRepo.On("New", ctx, input.ToDB()).Return(input.ToDB(), nil)
				linkCache.On("Set", ctx, input.ToDB()).Return(nil)
			},
			input:    input,
			callType: "new",
			err:      nil,
		},

		{
			name:          "empty origin",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {},
			input:         &entity.Link{CreatedAt: time.Now()},
			callType:      "new",
			err:           errors.New("origin is empty"),
		},

		{
			name: "new and store error",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkRepo.On("New", ctx, input.ToDB()).Return(nil, errors.New("store error"))
			},
			input:    input,
			callType: "new",
			err:      errors.New("store error"),
		},
		{
			name: "new and cache error",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkRepo.On("New", ctx, input.ToDB()).Return(input.ToDB(), nil)
				linkCache.On("Set", ctx, input.ToDB()).Return(errors.New("cache error"))
			},
			input:    input,
			callType: "new",
			err:      errors.New("cache error"),
		},

		{
			name: "get valid cached",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkCache.On("Get", ctx, input.Hash).Return(input.ToDB(), nil)
			},
			input:    input,
			callType: "get",
			err:      nil,
		},
		{
			name: "get and cache error",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkCache.On("Get", ctx, input.Hash).Return(nil, errors.New("cache error"))
			},
			input:    input,
			callType: "get",
			err:      errors.New("cache error"),
		},

		{
			name: "get uncached and store error",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkCache.On("Get", ctx, input.Hash).Return(nil, nil)
				linkRepo.On("Get", ctx, input.Hash).Return(nil, errors.New("store error"))
			},
			input:    input,
			callType: "get",
			err:      errors.New("store error"),
		},

		{
			name: "get empty result",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkCache.On("Get", ctx, input.Hash).Return(nil, nil)
				linkRepo.On("Get", ctx, input.Hash).Return(nil, nil)
			},
			input:    input,
			callType: "get",
			err:      nil,
		},

		{
			name: "link exists new cache error",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkRepo.On("Get", ctx, input.Hash).Return(input.ToDB(), nil)
				linkCache.On("Get", ctx, input.Hash).Return(nil, nil)
				linkCache.On("Set", ctx, input.ToDB()).Return(errors.New("cache error"))
			},
			input:    input,
			callType: "get",
			err:      errors.New("cache error"),
		},
		{
			name: "get uncached no errors",
			expectactions: func(ctx context.Context, linkRepo *mockDb.LinkRepository, linkCache *mockCache.LinkCache) {
				linkRepo.On("Get", ctx, input.Hash).Return(input.ToDB(), nil)
				linkCache.On("Get", ctx, input.Hash).Return(nil, nil)
				linkCache.On("Set", ctx, input.ToDB()).Return(nil)
			},
			input:    input,
			callType: "get",
		},
	}

	for _, test := range tests {
		t.Logf("running: %s", test.name)
		ctx := context.Background()
		linkRepo := &mockDb.LinkRepository{}
		linkCache := &mockCache.LinkCache{}
		usecase := NewLinkUsecase(linkRepo, linkCache)
		test.expectactions(ctx, linkRepo, linkCache)

		var err error
		if test.callType == "new" {
			_, err = usecase.New(ctx, test.input)
		} else if test.callType == "get" {
			_, err = usecase.Get(ctx, input.Hash)
		}

		if err != nil {
			if test.err != nil {
				//assert.Equal(t, test.err.Error(), err.Error())
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		}

		linkRepo.AssertExpectations(t)
		linkCache.AssertExpectations(t)
	}

}
