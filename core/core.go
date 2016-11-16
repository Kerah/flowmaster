package core

import (
	"net/url"
	"context"
)

type Message interface {
	Flow() string
	Body() []byte
	ContentType() uint32
	FlowType() int32
}

type Frame interface {
	Marshall() (data []byte, err error)
	Unmarshall(data []byte) error
}

type Flow interface {
	Id() string
	Name() string
	Headers() url.Values
	Context() context.Context
}

type Master interface {
	Read(name string, )
	Open(name string, flow interface{}) Master
	Sync() error
}

type Connector func(ctx context.Context) (Master, error)

