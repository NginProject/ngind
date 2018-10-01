package fetcher

import "github.com/NginProject/ngind/core/types"

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
