package statusHandler

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/core/check"
)

// AppStatusFacade will be used for handling multiple monitoring tools at once
type AppStatusFacade struct {
	handlers []core.AppStatusHandler
}

// NewAppStatusFacadeWithHandlers will receive the handlers which should receive monitored data
func NewAppStatusFacadeWithHandlers(aphs ...core.AppStatusHandler) (*AppStatusFacade, error) {
	if aphs == nil {
		return nil, ErrHandlersSliceIsNil
	}
	for _, aph := range aphs {
		if check.IfNil(aph) {
			return nil, ErrNilHandlerInSlice
		}
	}
	return &AppStatusFacade{
		handlers: aphs,
	}, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (asf *AppStatusFacade) IsInterfaceNil() bool {
	return asf == nil
}

// Increment method - will increment the value for a key for every handler
func (asf *AppStatusFacade) Increment(key string) {
	for _, ash := range asf.handlers {
		ash.Increment(key)
	}
}

// AddUint64 method - will increase the value for a key for every handler
func (asf *AppStatusFacade) AddUint64(key string, value uint64) {
	for _, ash := range asf.handlers {
		ash.AddUint64(key, value)
	}
}

// Decrement method - will decrement the value for a key for every handler
func (asf *AppStatusFacade) Decrement(key string) {
	for _, ash := range asf.handlers {
		ash.Decrement(key)
	}
}

// SetInt64Value method - will update the value for a key for every handler
func (asf *AppStatusFacade) SetInt64Value(key string, value int64) {
	for _, ash := range asf.handlers {
		ash.SetInt64Value(key, value)
	}
}

// SetUInt64Value method - will update the value for a key for every handler
func (asf *AppStatusFacade) SetUInt64Value(key string, value uint64) {
	for _, ash := range asf.handlers {
		ash.SetUInt64Value(key, value)
	}
}

// SetStringValue method - will update the value for a key for every handler
func (asf *AppStatusFacade) SetStringValue(key string, value string) {
	for _, ash := range asf.handlers {
		ash.SetStringValue(key, value)
	}
}

// Close method will close all the handlers
func (asf *AppStatusFacade) Close() {
	for _, ash := range asf.handlers {
		ash.Close()
	}
}
