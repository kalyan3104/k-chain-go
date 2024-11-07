package api

import "github.com/kalyan3104/k-chain-go/facade"

type noAPIInterface struct{}

// NewNoApiInterface will create a new instance of noAPIInterface
func NewNoApiInterface() *noAPIInterface {
	return new(noAPIInterface)
}

// RestApiInterface will return the value for disable api interface
func (n noAPIInterface) RestApiInterface(_ uint32) string {
	return facade.DefaultRestPortOff
}
