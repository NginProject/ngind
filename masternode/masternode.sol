pragma solidity ^0.5.0;

contract MN {

    struct Node {
        uint256 balance;
        bool isActive;
    }

    address[] public addrList;
    uint public addrNum;
    mapping(address => Node) public nodeList;   // balances, indexed by addresses

    constructor() public {
        addrNum = 0;
    }
    
    // Aka register
    function deposit() public payable {
        // append the address to list
        require(nodeList[msg.sender].balance + msg.value >= nodeList[msg.sender].balance);
        require((msg.value + nodeList[msg.sender].balance) > 1 * 10 ** 18);
        if (nodeList[msg.sender].balance == 0){
            addrList.push(msg.sender);
            addrNum = addrNum + 1 ;
        }
        
        nodeList[msg.sender].balance += msg.value;

        if (checkDeposit(msg.sender, nodeList[msg.sender].balance)){
            nodeList[msg.sender].isActive = true;
        }else{
            nodeList[msg.sender].isActive = false;
        }
    }
    
    function () external payable {
        //deposit();
    }
    
    // check the amount whether reach the threhold when deposit
    function checkDeposit(address addr, uint256 amount) public view returns (bool) {
        require(block.number>=0);        
        // return = (amount + nodeList[addr].balance) > (circulatingSupply(block.number) / 50); 
        return (amount + nodeList[addr].balance) > 1 * 10 ** 18; // for Debug
    }

    // If anyone need withdraw, all ngin will be withdrawn. DO NOT DEL
    function withdraw() public {
        uint256 amount = nodeList[msg.sender].balance;
        nodeList[msg.sender].balance = 0;
        nodeList[msg.sender].isActive = false;
        msg.sender.transfer(amount);
    }


    function circulatingSupply(uint h) public pure returns (uint256) {
        uint ERA_LEN = 100000;
        uint256 MAX_BLOCK_REWARD = 10 * 10 ** 18;
        uint cur_era = h / ERA_LEN;
        uint256 cs = 0;

        cs = MAX_BLOCK_REWARD*ERA_LEN*cur_era;

        return cs;
    }

    function id2Node(uint id) public view returns (address, uint256, bool){
        checkDeposit(addrList[id], nodeList[addrList[id]].balance);
        return (addrList[id], nodeList[addrList[id]].balance, nodeList[addrList[id]].isActive);
    }
}
