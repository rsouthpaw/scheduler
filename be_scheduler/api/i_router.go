package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func getRouter() *mux.Router {

	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	apiV1Router.HandleFunc("/login", login).Methods(http.MethodPost)

	filesRouter := apiV1Router.PathPrefix("/task").Subrouter()
	filesRouter.HandleFunc("/schedule", scheduleTask).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	return router
}

/**
 * Middleware should be present which authorizes each request be fore passing to router
 */
