package core

import (
	"net/url"
	"context"
)

type FlowID string
type FlowName string


type Message interface {
	Flow() FlowID
	Body() []byte
}

type Frame interface {
	Encode()
	Decode()
}

type Flow interface {
	Id() FlowID
	Name() FlowName
	Headers() url.Values
	Context() context.Context
}

type Master interface {
	Read(name FlowName, )
	Open(name string, flow interface{}) Master
	Sync() error
}

type Connector func(ctx context.Context) (Master, error)

