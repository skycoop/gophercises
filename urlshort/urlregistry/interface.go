package urlregistry

import "github.com/rs/xid"

type UrlRegistry interface {
	Register(url string) (xid.ID, error)
	Lookup(guid xid.ID) (string, error)
}
