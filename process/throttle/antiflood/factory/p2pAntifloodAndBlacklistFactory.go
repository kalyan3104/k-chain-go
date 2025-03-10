package factory

import (
	"context"
	"fmt"
	"time"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/kalyan3104/k-chain-go/config"
	antifloodDebug "github.com/kalyan3104/k-chain-go/debug/antiflood"
	"github.com/kalyan3104/k-chain-go/p2p"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/process/throttle/antiflood"
	"github.com/kalyan3104/k-chain-go/process/throttle/antiflood/blackList"
	"github.com/kalyan3104/k-chain-go/process/throttle/antiflood/disabled"
	"github.com/kalyan3104/k-chain-go/process/throttle/antiflood/floodPreventers"
	"github.com/kalyan3104/k-chain-go/statusHandler/p2pQuota"
	"github.com/kalyan3104/k-chain-go/storage/cache"
	storageFactory "github.com/kalyan3104/k-chain-go/storage/factory"
	"github.com/kalyan3104/k-chain-go/storage/storageunit"
	logger "github.com/kalyan3104/k-chain-logger-go"
)

var log = logger.GetOrCreate("p2p/antiflood/factory")

const defaultSpan = 300 * time.Second
const fastReactingIdentifier = "fast_reacting"
const slowReactingIdentifier = "slow_reacting"
const outOfSpecsIdentifier = "out_of_specs"
const outputIdentifier = "output"

var durationSweepP2PBlacklist = time.Second * 5

// AntiFloodComponents holds the handlers for the anti-flood and blacklist mechanisms
type AntiFloodComponents struct {
	AntiFloodHandler process.P2PAntifloodHandler
	BlacklistHandler process.PeerBlackListCacher
	FloodPreventers  []process.FloodPreventer
	TopicPreventer   process.TopicFloodPreventer
	PubKeysCacher    process.TimeCacher
}

// NewP2PAntiFloodComponents will return instances of antiflood and blacklist, based on the config
func NewP2PAntiFloodComponents(ctx context.Context, config config.Config, statusHandler core.AppStatusHandler, currentPid core.PeerID) (*AntiFloodComponents, error) {
	if check.IfNil(statusHandler) {
		return nil, p2p.ErrNilStatusHandler
	}
	if config.Antiflood.Enabled {
		return initP2PAntiFloodComponents(ctx, config, statusHandler, currentPid)
	}

	return &AntiFloodComponents{
		AntiFloodHandler: &disabled.AntiFlood{},
		BlacklistHandler: &disabled.PeerBlacklistCacher{},
		FloodPreventers:  make([]process.FloodPreventer, 0),
		TopicPreventer:   disabled.NewNilTopicFloodPreventer(),
		PubKeysCacher:    &disabled.TimeCache{},
	}, nil
}

func initP2PAntiFloodComponents(
	ctx context.Context,
	mainConfig config.Config,
	statusHandler core.AppStatusHandler,
	currentPid core.PeerID,
) (*AntiFloodComponents, error) {
	timeCache := cache.NewTimeCache(defaultSpan)
	p2pPeerBlackList, err := cache.NewPeerTimeCache(timeCache)
	if err != nil {
		return nil, err
	}

	publicKeysCache := cache.NewTimeCache(defaultSpan)

	fastReactingFloodPreventer, err := createFloodPreventer(
		ctx,
		mainConfig.Antiflood.FastReacting,
		mainConfig.Antiflood.Cache,
		statusHandler,
		fastReactingIdentifier,
		p2pPeerBlackList,
		currentPid,
	)
	if err != nil {
		return nil, fmt.Errorf("%w when creating fast reacting flood preventer", err)
	}

	slowReactingFloodPreventer, err := createFloodPreventer(
		ctx,
		mainConfig.Antiflood.SlowReacting,
		mainConfig.Antiflood.Cache,
		statusHandler,
		slowReactingIdentifier,
		p2pPeerBlackList,
		currentPid,
	)
	if err != nil {
		return nil, fmt.Errorf("%w when creating fast reacting flood preventer", err)
	}

	outOfSpecsFloodPreventer, err := createFloodPreventer(
		ctx,
		mainConfig.Antiflood.OutOfSpecs,
		mainConfig.Antiflood.Cache,
		statusHandler,
		outOfSpecsIdentifier,
		p2pPeerBlackList,
		currentPid,
	)
	if err != nil {
		return nil, fmt.Errorf("%w when creating out of specs flood preventer", err)
	}

	topicFloodPreventer, err := floodPreventers.NewTopicFloodPreventer(mainConfig.Antiflood.Topic.DefaultMaxMessagesPerSec)
	if err != nil {
		return nil, err
	}

	topicMaxMessages := mainConfig.Antiflood.Topic.MaxMessages
	setMaxMessages(topicFloodPreventer, topicMaxMessages)

	p2pAntiflood, err := antiflood.NewP2PAntiflood(
		p2pPeerBlackList,
		topicFloodPreventer,
		fastReactingFloodPreventer,
		slowReactingFloodPreventer,
		outOfSpecsFloodPreventer,
	)
	if err != nil {
		return nil, err
	}

	if mainConfig.Debug.Antiflood.Enabled {
		debugger, errDebugger := antifloodDebug.NewAntifloodDebugger(mainConfig.Debug.Antiflood)
		if errDebugger != nil {
			return nil, errDebugger
		}

		err = p2pAntiflood.SetDebugger(debugger)
		if err != nil {
			return nil, err
		}
	}

	startResettingTopicFloodPreventer(ctx, topicFloodPreventer, topicMaxMessages)
	startSweepingTimeCaches(ctx, p2pPeerBlackList, publicKeysCache)

	return &AntiFloodComponents{
		AntiFloodHandler: p2pAntiflood,
		BlacklistHandler: p2pPeerBlackList,
		PubKeysCacher:    publicKeysCache,
		FloodPreventers: []process.FloodPreventer{
			fastReactingFloodPreventer,
			slowReactingFloodPreventer,
			outOfSpecsFloodPreventer,
		},
		TopicPreventer: topicFloodPreventer,
	}, nil
}

func setMaxMessages(topicFloodPreventer process.TopicFloodPreventer, topicMaxMessages []config.TopicMaxMessagesConfig) {
	for _, topicMaxMsg := range topicMaxMessages {
		topicFloodPreventer.SetMaxMessagesForTopic(topicMaxMsg.Topic, topicMaxMsg.NumMessagesPerSec)
	}
}

func startResettingTopicFloodPreventer(
	ctx context.Context,
	topicFloodPreventer process.TopicFloodPreventer,
	topicMaxMessages []config.TopicMaxMessagesConfig,
	floodPreventers ...process.FloodPreventer,
) {
	localTopicMaxMessages := make([]config.TopicMaxMessagesConfig, len(topicMaxMessages))
	copy(localTopicMaxMessages, topicMaxMessages)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Debug("startResettingFloodPreventers's go routine is stopping...")
				return
			case <-time.After(time.Second):
			}

			for _, fp := range floodPreventers {
				fp.Reset()
			}
			for _, topicMaxMsg := range localTopicMaxMessages {
				topicFloodPreventer.ResetForTopic(topicMaxMsg.Topic)
			}
			topicFloodPreventer.ResetForNotRegisteredTopics()
		}
	}()
}

func startSweepingTimeCaches(ctx context.Context, p2pPeerBlackList process.PeerBlackListCacher, publicKeysCache process.TimeCacher) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Debug("startSweepingP2PPeerBlackList's go routine is stopping...")
				return
			case <-time.After(durationSweepP2PBlacklist):
			}

			p2pPeerBlackList.Sweep()
			publicKeysCache.Sweep()
		}
	}()
}

func createFloodPreventer(
	ctx context.Context,
	floodPreventerConfig config.FloodPreventerConfig,
	antifloodCacheConfig config.CacheConfig,
	statusHandler core.AppStatusHandler,
	quotaIdentifier string,
	blackListHandler process.PeerBlackListCacher,
	selfPid core.PeerID,
) (process.FloodPreventer, error) {
	cacheConfig := storageFactory.GetCacherFromConfig(antifloodCacheConfig)
	blackListCache, err := storageunit.NewCache(cacheConfig)
	if err != nil {
		return nil, err
	}

	blackListProcessor, err := blackList.NewP2PBlackListProcessor(
		blackListCache,
		blackListHandler,
		floodPreventerConfig.BlackList.ThresholdNumMessagesPerInterval,
		floodPreventerConfig.BlackList.ThresholdSizePerInterval,
		floodPreventerConfig.BlackList.NumFloodingRounds,
		time.Duration(floodPreventerConfig.BlackList.PeerBanDurationInSeconds)*time.Second,
		quotaIdentifier,
		selfPid,
	)
	if err != nil {
		return nil, err
	}

	antifloodCache, err := storageunit.NewCache(cacheConfig)
	if err != nil {
		return nil, err
	}

	quotaProcessor, err := p2pQuota.NewP2PQuotaProcessor(statusHandler, quotaIdentifier)
	if err != nil {
		return nil, err
	}

	basePeerMaxMessagesPerInterval := floodPreventerConfig.PeerMaxInput.BaseMessagesPerInterval
	peerMaxTotalSizePerInterval := floodPreventerConfig.PeerMaxInput.TotalSizePerInterval
	reservedPercent := floodPreventerConfig.ReservedPercent

	argFloodPreventer := floodPreventers.ArgQuotaFloodPreventer{
		Name:                      quotaIdentifier,
		Cacher:                    antifloodCache,
		StatusHandlers:            []floodPreventers.QuotaStatusHandler{quotaProcessor, blackListProcessor},
		BaseMaxNumMessagesPerPeer: basePeerMaxMessagesPerInterval,
		MaxTotalSizePerPeer:       peerMaxTotalSizePerInterval,
		PercentReserved:           reservedPercent,
		IncreaseThreshold:         floodPreventerConfig.PeerMaxInput.IncreaseFactor.Threshold,
		IncreaseFactor:            floodPreventerConfig.PeerMaxInput.IncreaseFactor.Factor,
	}
	floodPreventer, err := floodPreventers.NewQuotaFloodPreventer(argFloodPreventer)
	if err != nil {
		return nil, err
	}

	log.Debug("started antiflood & blacklist component",
		"type", quotaIdentifier,
		"interval in seconds", floodPreventerConfig.IntervalInSeconds,
		"base peerMaxMessagesPerInterval", basePeerMaxMessagesPerInterval,
		"peerMaxTotalSizePerInterval", core.ConvertBytes(peerMaxTotalSizePerInterval),
		"peerBanDurationInSeconds", floodPreventerConfig.BlackList.PeerBanDurationInSeconds,
		"thresholdNumMessagesPerSecond", floodPreventerConfig.BlackList.ThresholdNumMessagesPerInterval,
		"thresholdSizePerSecond", floodPreventerConfig.BlackList.ThresholdSizePerInterval,
		"numFloodingRounds", floodPreventerConfig.BlackList.NumFloodingRounds,
		"increase threshold", floodPreventerConfig.PeerMaxInput.IncreaseFactor.Threshold,
		"increase factor", floodPreventerConfig.PeerMaxInput.IncreaseFactor.Factor,
	)

	go func() {
		wait := time.Duration(floodPreventerConfig.IntervalInSeconds) * time.Second

		for {
			select {
			case <-ctx.Done():
				log.Debug("floodPreventer.Reset go routine is stopping...")
				return
			case <-time.After(wait):
			}

			floodPreventer.Reset()
		}
	}()

	return floodPreventer, nil
}
