package masternode

import (
	"github.com/NginProject/ngind/common"
	"net"
)

//types

type MasterNode struct {
	Addr     common.Address
	Coinbase common.Address
	Ip       net.IP
	Port     uint8
}
