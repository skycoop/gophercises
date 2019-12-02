package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/skycoop/gophercises/urlshort/urlregistry"
	"net/http"
	"net/url"
)

type RegisterHandler struct {
	registry urlregistry.UrlRegistry
}

type request struct {
	Url string `json:"url"`
}

type response struct {
	Guid string `json:"guid"`
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.WithField("error", err).Debug("Received invalid JSON")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	parsedUrl, err := url.Parse(req.Url)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   req.Url,
		}).Debug("Error when attempting to parse the URL")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if !parsedUrl.IsAbs() {
		log.WithFields(log.Fields{
			"error": err,
			"url":   parsedUrl,
		}).Debug("Provided url is not absolute")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	switch parsedUrl.Scheme {
	case "http":
		break
	case "https":
		break
	default:
		log.WithFields(log.Fields{
			"error": err,
			"url":   parsedUrl,
		}).Debug("Provided url has an invalid scheme")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	guid, err := h.registry.Register(parsedUrl.String())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   req.Url,
		}).Error("Error while registering new URL")
		panic(http.ErrAbortHandler)
	}

	res, err := json.Marshal(response{Guid: guid.String()})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"guid":  guid,
		}).Error("Error while marshalling JSON")
		panic(http.ErrAbortHandler)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"response": res,
		}).Error("Failed to write response")
		panic(http.ErrAbortHandler)
	}
}

func NewRegisterHandler(r urlregistry.UrlRegistry) *RegisterHandler {
	return &RegisterHandler{registry: r}
}
