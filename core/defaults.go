package core

import (
	"github.com/NginProject/ngind/logger/glog"
)

var (
	DefaultConfigMainnet *SufficientChainConfig
	DefaultConfigTestnet *SufficientChainConfig
)

func init() {

	var err error

	DefaultConfigMainnet, err = parseExternalChainConfig("/core/config/mainnet.json", assetsOpen)
	if err != nil {
		glog.Fatal("Error parsing mainnet defaults from JSON:", err)
	}
	DefaultConfigTestnet, err = parseExternalChainConfig("/core/config/testnet.json", assetsOpen)
	if err != nil {
		glog.Fatal("Error parsing testnet defaults from JSON:", err)
	}
}
