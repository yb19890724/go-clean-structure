package main

import (
	"fmt"
	"github.com/yb19890724/go-clean-structure/pkg/product/delivery/http/rest"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/listing"
	"github.com/yb19890724/go-clean-structure/pkg/product/storges/json"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service
	var lister listing.Service

	s, _ := json.NewStorage()

	// 注入存储库
	adder = adding.NewService(s)
	lister = listing.NewService(s)

	router := rest.Handler(adder, lister)

	fmt.Println("The product server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
