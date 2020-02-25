package main

import (
	"astro_pro/api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("astro_pro.conf")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func makeRouter() *mux.Router {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir(viper.GetString("api.ui_path")))

	router.Handle("/playground", handler.Playground("GraphQL playground", "/query")).Methods("GET")
	router.Handle("/query", handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}}))).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(viper.GetString("api.ui_path"), "index.html")
		body, _ := ioutil.ReadFile(path)

		fmt.Fprint(w, string(body))
	}))

	return router
}

func main() {
	port := viper.GetString("api.port")

	router := makeRouter()

	log.Printf("connect to http://localhost:%s/ for ASTRO Professional", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
