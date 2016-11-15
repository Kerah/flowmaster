package connectors

import (
	"context"
	"github.com/Kerah/flowmaster/consts"
)

func GetServiceName(ctx context.Context) (serviceName string) {
	serviceName, _ = ctx.Value(consts.RemoteServiceName).(string)
	return
}

func SetServiceName(ctx context.Context, serviceName string) context.Context {
	context.WithValue(ctx, consts.RemoteServiceName, serviceName)
}
