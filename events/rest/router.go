package rest

import (
	"github.com/gorilla/mux"
	"github.com/ricardo-gomez/reservations/lib/persistence"
	"net/http"
)

func ServerAPI(endpoint, endpointTLS string, dbHandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := newEventHandler(dbHandler)
	router := mux.NewRouter()
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventsHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	httpErrChan := make(chan error)
	httpTlsErrChan := make(chan error)

	go func() {httpErrChan <- http.ListenAndServe(endpoint, router)}()
	go func() {httpTlsErrChan <- http.ListenAndServeTLS(endpointTLS, "cert.pem", "key.pem", router)}()
	return httpErrChan, httpTlsErrChan
}
