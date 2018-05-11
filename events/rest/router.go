package rest

import (
	"github.com/gorilla/mux"
	"github.com/ricardo-gomez/reservations/lib/persistence"
	"net/http"
)

func ServerAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := newEventHandler(dbHandler)
	router := mux.NewRouter()
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventsHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, router)
}
