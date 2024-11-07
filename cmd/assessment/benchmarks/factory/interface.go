package factory

import "github.com/kalyan3104/k-chain-go/cmd/assessment/benchmarks"

type benchmarkCoordinator interface {
	RunAllTests() *benchmarks.TestResults
	IsInterfaceNil() bool
}
