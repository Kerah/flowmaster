package flowmaster

import (
	"context"
	"github.com/Kerah/flowmaster/core"
	"github.com/Kerah/flowmaster/consts"
	"github.com/Kerah/flowmaster/connectors"
	"time"
)

func defaultConnector(ctx context.Context) (master core.Master, err error) {
	return
}

func createDefaultRequestContext(serviceName string) context.Context {
	ctx := context.Background()
	ctx = connectors.SetServiceName(ctx, serviceName)
	return ctx

}

func Open(serviceName string, ctx context.Context) (master core.Master, err error) {
	var connector core.Connector = defaultConnector
	if cn := ctx.Value(consts.ConnectorFn); cn !=nil {
		if fnc, ok := cn.(core.Connector); ok {
			connector = fnc
		}
	}
	var req context.Context

	if ctx == nil {
		req = createDefaultRequestContext(serviceName)
	} else {
		req = connectors.SetServiceName(ctx, serviceName)
	}
	return connector(connectors.SetServiceName(req, serviceName))
}

