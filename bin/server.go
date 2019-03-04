package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/michalnov/psychology/bin/core"
	"github.com/michalnov/psychology/bin/core/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var r *mux.Router

func main() {
	fmt.Println("Hello server")
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Failed to read config")
		return
	}
	var c core.Core
	err = c.LoadTests()
	if err != nil {
		fmt.Println("failed to load tests")
		return
	}
	r = mux.NewRouter()
	r.HandleFunc("/register", notImplemented).Methods("post")
	r.HandleFunc("/gettest", notImplemented).Methods("post")
	r.HandleFunc("/starttest", notImplemented).Methods("post")
	r.HandleFunc("/answer", notImplemented).Methods("post")
	r.HandleFunc("/finishtest", notImplemented).Methods("post")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	r.HandleFunc("/hello", notImplemented)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(conf.Server.Port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("not implemented yet"))
}
