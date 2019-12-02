package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/skycoop/gophercises/cyoa"
	"html/template"
	"net/http"
)

type ArcHandler struct {
	story *cyoa.Story
	tmpl  *template.Template
}

func (h ArcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.WithField("arc", vars["arc"]).Debug("Serving arc page")
	arc, err := h.story.GetArc(vars["arc"])
	if err != nil {
		log.WithFields(log.Fields{"error": err, "arc": vars["arc"]}).Debug("Failed to get arc")
		http.NotFound(w, r)
		panic(http.ErrAbortHandler)
	}

	w.Header().Set("Content-Type", "text/html")
	err = h.tmpl.ExecuteTemplate(w, "arc.gohtml", arc)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "arc": vars["arc"]}).Error("Failed to template arc page")
		panic(http.ErrAbortHandler)
	}
}

func NewArcHandler(s *cyoa.Story) (*ArcHandler, error) {
	log.Debug("Loading templates")
	t, err := template.ParseFiles("cyoa/web/templates/arc.gohtml")
	if err != nil {
		return nil, fmt.Errorf("failed to parse template file: %w", err)
	}

	return &ArcHandler{story: s, tmpl: t}, nil
}
