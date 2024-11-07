package chronology

import (
	"time"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/consensus"
	"github.com/kalyan3104/k-chain-go/ntp"
)

// ArgChronology holds all dependencies required by the chronology component
type ArgChronology struct {
	GenesisTime      time.Time
	RoundHandler     consensus.RoundHandler
	SyncTimer        ntp.SyncTimer
	Watchdog         core.WatchdogTimer
	AppStatusHandler core.AppStatusHandler
}
