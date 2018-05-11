package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ricardo-gomez/reservations/lib/persistence"
	"net/http"
	"strings"
)

type eventServiceHandler struct {
	dbHandler persistence.DatabaseHandler
}

func newEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbHandler: databaseHandler,
	}
}

func (eventHandler *eventServiceHandler) findEventHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		res.WriteHeader(400)
		fmt.Fprint(res, "{error: no hay criterio de busqueda, la url debe ser /id/222, /name/evento}")
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		res.WriteHeader(400)
		fmt.Fprint(res, "{error: falta parametro para buscar, la url debe ser /id/222, /name/evento}")
		return
	}
	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eventHandler.dbHandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eventHandler.dbHandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(res, "{error: %s}", err)
		return
	}
	res.Header().Set("Content-Type", "Aplication/json;charset=utf8")
	json.NewEncoder(res).Encode(&event)
}

func (eventHandler *eventServiceHandler) allEventsHandler(res http.ResponseWriter, req *http.Request) {
	events, err := eventHandler.dbHandler.FindAllAvailableEvents()
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprintf(res, "{error: Error %s}", err)
		return
	}
	res.Header().Set("Content-Type", "aplication/json;charset=utf8")
	err = json.NewEncoder(res).Encode(&events)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprintf(res, "{error: ocurrio un error %s}", err)
	}
}

func (eventHandler *eventServiceHandler) newEventHandler(res http.ResponseWriter, req *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprintf(res, "{error: error al decodificar el request}", err)
		return
	}
	id, err := eventHandler.dbHandler.AddEvent(event)
	if nil != err {
		res.WriteHeader(500)
		fmt.Fprintf(res, "{error: error al guardar evento %d %s}", id, err)
		return
	}
}
