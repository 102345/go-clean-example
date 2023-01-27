package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/di"
	"github.com/marc/go-clean-example/infra-structure/middlewares/authentication"
	"github.com/spf13/viper"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/marc/go-clean-example/adapter/http/docs"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Marcilio Gomes
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /

func main() {

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	router := configureRouters(conn)

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

func configureRouters(conn postgres.PoolInterface) *mux.Router {

	postgres.RunMigrations()

	productService := di.ConfigProductDI(conn)
	userService := di.ConfigUserDI(conn)

	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	router.Handle("/product",
		http.HandlerFunc(authentication.Logger((authentication.Authenticate(productService.Create, false))))).Methods("POST")

	router.Handle("/product/{product_id}",
		http.HandlerFunc(authentication.Logger((authentication.Authenticate(productService.Update, false))))).Methods("PUT")

	router.Handle("/product/{product_id}",
		http.HandlerFunc(authentication.Logger((authentication.Authenticate(productService.Delete, true))))).Methods("DELETE")

	router.Handle("/product",
		http.HandlerFunc(authentication.Logger((authentication.Authenticate(productService.Fetch, false))))).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	router.Handle("/user",
		http.HandlerFunc(authentication.Logger((authentication.Authenticate(userService.Create, false))))).Methods("POST")

	return router
}
