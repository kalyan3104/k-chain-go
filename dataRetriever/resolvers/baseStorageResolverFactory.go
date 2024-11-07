package resolvers

import "github.com/kalyan3104/k-chain-go/storage"

func createBaseStorageResolver(
	storer storage.Storer,
	isFullHistoryNode bool,
) baseStorageResolver {
	if isFullHistoryNode {
		return &baseFullHistoryResolver{
			storer: storer,
		}
	}

	return &baseSimpleResolver{
		storer: storer,
	}
}
