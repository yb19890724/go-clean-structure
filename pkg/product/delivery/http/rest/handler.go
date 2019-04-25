package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"net/http"
)

func Handler(a adding.Service,) http.Handler {
	router := httprouter.New()

	router.GET("/products", products(a))
	/*router.GET("/product/:id", product())

	router.POST("/store/:id", store())*/

	return router
}

func products(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		var newProduct adding.Product
		err := decoder.Decode(&newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.Add(newProduct)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New beer added.")
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