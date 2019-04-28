package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yb19890724/go-clean-structure/pkg/product/delivery/http/rest"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/adding"
	"github.com/yb19890724/go-clean-structure/pkg/product/service/listing"
	"github.com/yb19890724/go-clean-structure/pkg/product/storges/json"
	"log"
	"net/http"
	"os"
)

const ConfigName = "config"

func init() {

	viper.SetEnvPrefix(ConfigName)
	viper.AddConfigPath("$GOPATH/src/github.com/yb19890724/go-clean-structure/configs")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", ConfigName, err))
		os.Exit(1)
	}

}

func main() {

	var adder adding.Service
	var lister listing.Service

	s, _ := json.NewStorage()

	// 注入存储库
	adder = adding.NewService(s)
	lister = listing.NewService(s)

	router := rest.Handler(adder, lister)

	fmt.Printf("The product server is on tap now: http://localhost %s",viper.GetString("servers.port"))
	log.Fatal(http.ListenAndServe(viper.GetString("servers.port"), router))
}
