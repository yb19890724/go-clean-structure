package json

import (
	"github.com/nanobox-io/golang-scribble"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"math/rand"
	"path"
	"runtime"
	"strconv"
	"time"
)

const (
	dir = "/data/" // 数据存储目录

	CollectionProduct = "products"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) Add(b adding.Product) error {

	newP := adding.Product{
		ID:          rand.Int(),
		Name:        b.Name,
		Price:       rand.Float32(),
		Description: b.Description,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	resource := strconv.Itoa(newP.ID)
	if err := s.db.Write(CollectionProduct, resource, newP); err != nil {
		return err
	}
	return nil
}
