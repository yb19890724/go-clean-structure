package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/listing"
	response "github.com/yb19890724/go-clean-structure/tools/response/json"
	"net/http"
	"strconv"
)

func Handler(a adding.Service, l listing.Service) http.Handler {

	router := httprouter.New()

	router.POST("/products", store(a))

	router.GET("/products", products(l))

	router.GET("/product/:id", product(l))

	return router
}

// 存储产品数据 action
func store(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		decoder := json.NewDecoder(r.Body)

		var newProduct adding.Product

		if err := decoder.Decode(&newProduct);err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		s.Add(newProduct)

		response.WithCreated(w)

	}

}

// 获取产品列表
func products(l listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		list := l.Products()

		response.ResponseJson(w, list)

	}

}

// 获取指定产品
func product(l listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		ID, err := strconv.Atoi(p.ByName("id"))

		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid product ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}


		product,err := l.Product(ID)

		if err == listing.ErrNotFound {
			http.Error(w, "The product you requested does not exist.", http.StatusNotFound)
			return
		}

		response.ResponseJson(w, product)
	}
}
