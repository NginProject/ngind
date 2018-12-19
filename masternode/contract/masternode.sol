pragma solidity ^0.5.0;

contract Masternode {

    uint256 constant MAX_BLOCK_REWARD = 10 * 10 ** 18;
    uint constant ERA_LEN = 100000;
    uint constant COEFFICIENT_UP = 249;
    uint constant COEFFICIENT_DOWN = 250;

    struct Node {
        uint256 balance;
        bool isActive;
    }

    address[] public itList;
    mapping(address => Node) public nodeList;   // balances, indexed by addresses

    event Active(address _account);
    event Deactive(address _account);

    // Aka register
    function deposit(uint256 amount) public payable {
        require(msg.value == amount, "Error: wrong deposit amount");
        // append the address to list
        bool notInItList = true;
        for (uint i = 0; i < itList.length; i++) {
            if (itList[i] == msg.sender){
                notInItList = false;
                break;
            }
        }
        if (notInItList){
            itList.push(msg.sender);
        }
        // add the node and check master or not
        nodeList[msg.sender].balance += amount;
        if (checkDeposit(msg.sender, amount, block.number)){
            // more than threhold
            nodeList[msg.sender].isActive = true;
        }else{
            // not reach
            nodeList[msg.sender].isActive = false;
        }
    }

    // If anyone need withdraw, all ngin will be withdrawn. DO NOT DEL
    function withdraw() public {
        uint256 amount = nodeList[msg.sender].balance;
        msg.sender.transfer(amount);
        for (uint i = 0; i < itList.length; i++) {
            if (itList[i] == msg.sender){
                delete itList[i];
            }
        }
        delete nodeList[msg.sender];
    }

    function updateStatus() public {
        for (uint i = 0; i < itList.length; i++) {
            if (checkDeposit(itList[i], itList[i].balance, block.number)){
                // more than threhold
                nodeList[itList[i]].isActive = true;
            }else{
                // not reach
                nodeList[itList[i]].isActive = false;
            }
        }

    }

    function() external payable {
        updateStatus();
    }

    // check the amount whether reach the threhold when deposit
    function checkDeposit(address account, uint256 amount, uint256 h) public view returns (bool) {
        bool rtn = (amount + nodeList[account].balance) > (circulatingSupply(h) / 50); // > 1 * 10 ** 18; // for Debug
        return rtn;
    }

    function circulatingSupply(uint h) public pure returns (uint256) {
        uint cur_era = h / ERA_LEN;
        uint256 cs = 0;

        for (uint era = 0; era <= cur_era; era++){
            if (era == cur_era){
                uint256 r = MAX_BLOCK_REWARD * (COEFFICIENT_UP**era / COEFFICIENT_DOWN**era);
                cs = cs + r * (h % (cur_era*ERA_LEN));
                continue;
            }else{
                uint256 r = MAX_BLOCK_REWARD * (COEFFICIENT_UP**era / COEFFICIENT_DOWN**era);
                cs = cs + r * ERA_LEN;
                continue;
            }
        }
        return cs;
    }

}
