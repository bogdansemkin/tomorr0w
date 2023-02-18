package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type API struct {
	router *chi.Mux
}

type Card struct {
	Number string `json:"number"`
}

const cardNumber = "3333 4444 5555 6666"

func main() {
	apiServer, err := New().Setup()
	if err != nil {
		log.Printf("error during setup server, %s", err)
	}

	err = apiServer.Serve()
	if err != nil {
		log.Printf("unexpected shutdown")
	}

}

func New() *API {
	router := chi.NewRouter()

	return &API{router: router}
}

func (a *API) Setup() (*API, error) {
	a.router.Get("/api/v1/card/number", CardHandler)

	return a, nil
}

func (a *API) Serve() error {
	srv := &http.Server{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Addr:         ":8080",
		Handler:      a.router,
	}

	return srv.ListenAndServe()
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Card{Number: cardNumber})
	if err != nil {
		log.Printf("error during marshaling json, %s", err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Printf("error during writing json, %s", err)
	}
}
