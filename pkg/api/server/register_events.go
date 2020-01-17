package server

import (
	"net/http"

	"github.com/containers/libpod/pkg/api/handlers"
	"github.com/gorilla/mux"
)

func (s *APIServer) RegisterEventsHandlers(r *mux.Router) error {
	// == Returns events filtered on query parameters
	//
	// ?since=string   : start streaming events from this time
	// ?until=string   : stop streaming events later than this
	// ?filters=string : JSON-encoded map[string][]string of constraints
	//
	// 200 OK
	// 500 Failed  #InternalError
	r.Handle(VersionedPath("/events"), APIHandler(s.Context, handlers.GetEvents)).Methods(http.MethodGet)
	return nil
}
