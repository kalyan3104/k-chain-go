package antiflood

import "github.com/kalyan3104/k-chain-go/process"

func (af *p2pAntiflood) Debugger() process.AntifloodDebugger {
	return af.debugger
}
