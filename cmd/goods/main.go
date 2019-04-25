package main

import (
	"fmt"
	"github.com/yb19890724/go-clean-structure/pkg/product/delivery/http/rest"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"github.com/yb19890724/go-clean-structure/pkg/product/storges/json"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service

	s, _ := json.NewStorage()

	// 注入存储库
	adder = adding.NewService(s)
	router := rest.Handler(adder)

	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
