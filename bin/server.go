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
	fmt.Println(conf.Server.Port)
	if err != nil {
		fmt.Println("Failed to read config")
		return
	}
	var c core.Core
	master, slave, err := config.InitializeDb()
	if err != nil {
		fmt.Println("db init fail")
		return
	}
	c.DbMaster = master
	c.DbSlave = slave
	err = c.LoadTests()
	if err != nil {
		fmt.Println("failed to load tests")
		return
	}
	r = mux.NewRouter()
	r.HandleFunc("/user", c.UserHandler)
	r.HandleFunc("/gettest", c.GetTest).Methods("POST")
	r.HandleFunc("/ping", c.Ping).Methods("GET")
	r.HandleFunc("/answer", c.Answerque).Methods("POST")
	r.HandleFunc("/finishtest", notImplemented).Methods("POST")
	r.HandleFunc("/gas", c.GetGas).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(conf.Server.Port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("not implemented yet"))
}
