package memory

import (
	"sync"
	"time"

	"github.com/Neko2h/shortener/internal/entity"
)

const Timeout = 5

var InmemoryStorage = make(map[int64]*entity.LinkDB)

type Memory struct {
	sync.Mutex
}

func NewInMemoryDb() *Memory {
	return &Memory{}
}

func sliceSearch(slice map[int64]*entity.LinkDB, needle string) *entity.LinkDB {
	for _, element := range slice {
		if element.Hash == needle {
			return element
		}
	}
	return nil
}

func addElementToStorage(newLink *entity.LinkDB) *entity.LinkDB {
	newLink.ID = int64(len(InmemoryStorage) + 1)
	newLink.CreatedAt = time.Now()
	newLink.UpdatedAt = time.Now()

	InmemoryStorage[newLink.ID] = newLink
	return newLink
}
