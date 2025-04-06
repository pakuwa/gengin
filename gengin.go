package gengin

import "github.com/pakuwa/gengin/schema/route"

type Schema struct {
	Interface
}

type Interface interface {
	Routes() []Route
}

type Route interface {
	Descriptor() *route.Descriptor
}
