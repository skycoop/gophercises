package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"github.com/skycoop/gophercises/urlshort/urlregistry"
	"net/http"
)

type LookupHandler struct {
	registry urlregistry.UrlRegistry
}

func (h LookupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid, err := xid.FromString(vars["guid"])
	if err != nil {
		log.WithFields(log.Fields{
			"guid":  vars["guid"],
			"error": err,
		}).Debug("Received invalid guid")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var url string
	url, err = h.registry.Lookup(guid)
	if err != nil {
		log.WithFields(log.Fields{
			"guid":  guid,
			"error": err,
		}).Debug("Failed to lookup URL")
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func NewLookupHandler(r urlregistry.UrlRegistry) *LookupHandler {
	return &LookupHandler{registry: r}
}
