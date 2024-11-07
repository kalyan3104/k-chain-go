package factory

import (
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/debug/handler"
)

// NewInterceptorDebuggerFactory will instantiate an InterceptorDebugHandler based on the provided config
func NewInterceptorDebuggerFactory(config config.InterceptorResolverDebugConfig) (InterceptorDebugHandler, error) {
	if !config.Enabled {
		return handler.NewDisabledInterceptorDebugHandler(), nil
	}

	return handler.NewInterceptorDebugHandler(config)
}
