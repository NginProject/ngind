package masternode

import (
	"github.com/NginProject/ngind/accounts/abi/bind"
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/masternode/contract"
	"math/big"
)

//Test initial message gets set up correctly
func GetNodeList(mn *contract.MN) []common.Address {

	var NodeList []common.Address

	for i:=int64(0); i<=50 ; i++{
		opts := new(bind.CallOpts)
		//opts.BlockNumber = blockNumber
		addr, _, isActive, err := mn.Id2Node(opts, big.NewInt(i))

		if isActive == true{
			NodeList = append(NodeList,addr)
		}

		if err != nil{
			return NodeList
		}
	}

	return nil
}
