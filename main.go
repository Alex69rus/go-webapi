package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
)

func main() {
	// dbConfig := DbConfiguration{
	// 	Host:     "localhost",
	// 	User:     "postgres",
	// 	Password: "qwerty123",
	// 	Schema:   "golang",
	// }

	dbConfig := DbConfiguration{}
	readConfiguration("Db", &dbConfig)

	handler := NewHandler(&dbConfig)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.homeLink).Methods(http.MethodGet)
	router.HandleFunc("/book", handler.createBook).Methods(http.MethodPost)
	router.HandleFunc("/books", handler.getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.getBookByID).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.updateExistingBook).Methods(http.MethodPatch)
	router.HandleFunc("/books/{id}", handler.deleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":18132", router))
}

func readConfiguration(key string, cfg interface{}) {
	f, err := os.Open("config.json")
	processError(err)

	fullCfg := map[string]interface{}{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&fullCfg)
	processError(err)

	cfgByKey, ok := fullCfg[key]
	if !ok {
		processError(fmt.Errorf("Coud not found key: %v", key))
	}

	res, err := json.Marshal(cfgByKey)
	processError(err)

	json.Unmarshal(res, cfg)

	readEnv(cfg)
}

func readEnv(cfg interface{}) {
	err := envconfig.Process("", cfg)
	processError(err)
}

func processError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
