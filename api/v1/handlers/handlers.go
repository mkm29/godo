package handlers

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	// "github.com/mkm29/godo/api/types"
)

func Routes( /* any dependency injection comes here*/ ) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthz", Healthz)
	r.Get("/items", GetItems)
	r.Get("/item/:itemId", GetItem)
	r.Post("/item", CreateItem)
	r.Delete("/item/:itemId", DeleteItem)
	return r
}

func Healthz(rw http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, `{"alive": true}`)
}

// GetItems implementation
func GetItems(rw http.ResponseWriter, r *http.Request) {
	// get items from db

	// return list of items
}

// GetItem implements a simple handler to use the RW and Request object
func GetItem(rw http.ResponseWriter, r *http.Request) {
	// get item from db

	// return item
}

// CreateItems shows different ways you can pass dependencies into your handler
func CreateItem(rw http.ResponseWriter, r *http.Request) {
	// create new item

	// get POST data

	// save item

	// return JSON status
}

// DeleteItem implements a simple handler to use the RW and Request object
func DeleteItem(rw http.ResponseWriter, r *http.Request) {
	// get item from db

	// delete item
}
