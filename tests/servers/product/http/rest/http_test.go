package rest

import (
	"github.com/gavv/httpexpect"
	"github.com/yb19890724/go-clean-structure/tools/response/json"
	"net/http"
	"testing"
)

var testUrl string

func init() {

	testUrl ="http://localhost:8080"

}

// @test products
func TestProducts(t *testing.T) {

	e := httpexpect.New(t, testUrl)

	e.GET("/products").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("msg", "请求成功!").
		ValueEqual("code", json.Success)
}

// @test product
func TestProduct(t *testing.T) {

	e := httpexpect.New(t, testUrl)

	e.GET("/product/5577006791947779410").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("msg", "请求成功!").
		ValueEqual("code", json.Success)
}

// @test product
func TestStore(t *testing.T) {

	e := httpexpect.New(t, testUrl)

	p := map[string]interface{}{
		"Name":        "test product",
		"Description": "test description",
	}

	ct := "application/json;charset=utf-8"

	e.POST("/products").
		WithHeader("ContentType", ct).
		WithJSON(p).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object().
		ValueEqual("msg", "添加成功!").
		ValueEqual("code", json.Success)
}
