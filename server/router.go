package server

import (
	"RestApiServer"
	"github.com/gorilla/mux"
	"net/http"
)

func PrepareRoutingTable() http.Handler {
	router := mux.NewRouter()
	router.PathPrefix("/getUser").Methods(http.MethodGet).HandlerFunc(RestApiServer.GetUser)
	return router
}
