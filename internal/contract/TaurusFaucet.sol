// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/access/Ownable.sol";

contract TaurusFaucet is Ownable {

    event FaucetDripped(address _recipient,uint _amount);


    uint256 public ONE_DAY_SECONDS = 86400;
    mapping(address => uint256) public lastClaim;

    constructor() public Ownable(msg.sender) {
   }
 

    function drip(address _recipient, uint _amount) external onlyOwner {
        // Check if same githubid has claimed past 24hours
        require(canDrip(lastClaim[_recipient]), "Has claimed in the last 24hours");
        // Drip Ether
        (bool sent,) = _recipient.call{value: _amount}("");
        require(sent, "Failed dripping Axc");
        lastClaim[_recipient] = block.timestamp;

        emit FaucetDripped(_recipient,_amount);
    }
 
    function canDrip(uint256 _lastClaimTime) internal view returns (bool) {
        if(_lastClaimTime > block.timestamp)  {
            return false;
        }

        if(_lastClaimTime <= 0) {
            return true;
        }

        return ((block.timestamp - _lastClaimTime) >= ONE_DAY_SECONDS);
    }

    receive() external payable {}
}