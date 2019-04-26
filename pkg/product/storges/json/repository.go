package json

import (
	"encoding/json"
	"github.com/nanobox-io/golang-scribble"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/listing"
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

func (s *Storage) Add(p adding.Product) error {

	newP := Product{
		ID:          rand.Int(),
		Name:        p.Name,
		Price:       rand.Float32(),
		Description: p.Description,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	resource := strconv.Itoa(newP.ID)
	if err := s.db.Write(CollectionProduct, resource, newP); err != nil {
		return err
	}
	return nil
}


func (s *Storage) Product(id int) (listing.Product, error) {
	var p Product
	var product listing.Product

	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionProduct, resource, &p); err != nil {
		return product, listing.ErrNotFound
	}

	product.ID = p.ID
	product.Name = p.Name
	product.Price = p.Price
	product.Description = p.Description
	product.Created = p.Created
	product.Updated = p.Updated

	return product, nil
}

func (s *Storage) Products() []listing.Product {

	list := []listing.Product{}

	records, err := s.db.ReadAll(CollectionProduct)

	if err != nil {

		return list
	}

	for _, r := range records {
		var p Product
		var product listing.Product

		if err := json.Unmarshal([]byte(r), &p); err != nil {

			return list

		}

		product.ID = p.ID
		product.Name = p.Name
		product.Price = p.Price
		product.Description = p.Description
		product.Created = p.Created
		product.Updated = p.Updated

		list = append(list, product)
	}

	return list
}