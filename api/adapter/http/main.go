package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/102345/authenticationJWT/authenticationJWT"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/di"
	"github.com/spf13/viper"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/marc/go-clean-example/adapter/http/docs"
	"github.com/marc/go-clean-example/adapter/http/stockproductservice"
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

	router := configRouters(conn)

	headers, methods, origins := configCORS()

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), handlers.CORS(methods, origins, headers)(router))
}

func configCORS() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Origin", "X-Requested-With", "Accept"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "GET"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return headers, methods, origins
}

func configRouters(conn postgres.PoolInterface) *mux.Router {

	postgres.RunMigrations()

	productService := di.ConfigProductDI(conn)
	userService := di.ConfigUserDI(conn)

	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	router.Handle("/products",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(productService.Create, false))))).Methods("POST")

	router.Handle("/products/{product_id}",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(productService.Update, false))))).Methods("PUT")

	router.Handle("/products/{product_id}",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(productService.Delete, false))))).Methods("DELETE")

	router.Handle("/products",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(productService.Fetch, false))))).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	router.Handle("/products/{product_id}",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(productService.FindById, false))))).Methods("GET")

	router.Handle("/users",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(userService.Create, false))))).Methods("POST")
	router.Handle("/login",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(userService.Login, false))))).Methods("POST")
	router.Handle("/users/{user_id}",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(userService.Update, true))))).Methods("PUT")
	router.Handle("/users/{user_id}",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(userService.Delete, true))))).Methods("DELETE")
	router.Handle("/users",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(userService.Fetch, true))))).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	router.Handle("/stockproducts",
		http.HandlerFunc(authenticationJWT.Logger((authenticationJWT.Authenticate(stockproductservice.SendMessageStockProduct, false))))).Methods("POST")

	return router
}
