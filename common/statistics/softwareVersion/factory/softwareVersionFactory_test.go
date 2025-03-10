package factory

import (
	"testing"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/config"
	statusHandlerMock "github.com/kalyan3104/k-chain-go/testscommon/statusHandler"
	"github.com/stretchr/testify/assert"
)

func TestNewSoftwareVersionFactory_NilStatusHandlerShouldErr(t *testing.T) {
	t.Parallel()

	factory, err := NewSoftwareVersionFactory(nil, config.SoftwareVersionConfig{})

	assert.Equal(t, core.ErrNilAppStatusHandler, err)
	assert.Nil(t, factory)
}

func TestSoftwareVersionFactory_Create(t *testing.T) {
	t.Parallel()

	statusHandler := &statusHandlerMock.AppStatusHandlerStub{}
	factory, _ := NewSoftwareVersionFactory(statusHandler, config.SoftwareVersionConfig{PollingIntervalInMinutes: 1})
	softwareVersionChecker, err := factory.Create()

	assert.Nil(t, err)
	assert.NotNil(t, softwareVersionChecker)
}
