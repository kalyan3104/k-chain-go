package cutoff

import "github.com/kalyan3104/k-chain-go/config"

// CreateBlockProcessingCutoffHandler will create the desired block processing cutoff handler based on configuration
func CreateBlockProcessingCutoffHandler(cfg config.BlockProcessingCutoffConfig) (BlockProcessingCutoffHandler, error) {
	if !cfg.Enabled {
		return NewDisabledBlockProcessingCutoff(), nil
	}

	return NewBlockProcessingCutoffHandler(cfg)
}
