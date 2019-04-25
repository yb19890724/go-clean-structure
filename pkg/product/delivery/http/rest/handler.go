package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	response "github.com/yb19890724/go-clean-structure/tools/response/json"
	"net/http"
)

func Handler(a adding.Service,) http.Handler {

	router := httprouter.New()

	router.POST("/products", store(a))

	/*router.GET("/product/:id", product())

	router.POST("/store/:id", store())*/

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

/*func product() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_ := json.NewDecoder(r.Body)

	}
}

func store() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_ := json.NewDecoder(r.Body)

	}
}
*/