package api

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	return getRouter()
}
