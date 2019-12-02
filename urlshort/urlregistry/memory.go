package urlregistry

import (
	"errors"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"sync"
)

type MemoryUrlRegistry struct {
	mutex     sync.Mutex
	guidToUrl map[xid.ID]string
	urlToGuid map[string]xid.ID
}

func (r *MemoryUrlRegistry) Register(url string) (xid.ID, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	guid, exists := r.urlToGuid[url]
	if exists {
		return guid, nil
	}

	guid = xid.New()
	r.urlToGuid[url] = guid
	r.guidToUrl[guid] = url

	log.WithFields(log.Fields{
		"guid": guid,
		"url":  url,
	}).Debug("Registered new URL")
	return guid, nil
}

func (r *MemoryUrlRegistry) Lookup(guid xid.ID) (string, error) {
	url, exists := r.guidToUrl[guid]
	if exists {
		return url, nil
	} else {
		return "", errors.New("no url registered to guid")
	}
}

func NewMemoryUrlRegistry() *MemoryUrlRegistry {
	return &MemoryUrlRegistry{
		mutex:     sync.Mutex{},
		guidToUrl: make(map[xid.ID]string),
		urlToGuid: make(map[string]xid.ID),
	}
}
