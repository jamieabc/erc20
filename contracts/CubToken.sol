pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract CubToken is ERC20, Ownable{
  uint256 public totalSupply;
  string public name;
  string public symbol;
  uint32 public decimals;
  mapping(address => uint256) public balances;

  constructor() public {
    symbol = "CUB";
    name = "CubToken";
    decimals = 5;
    totalSupply = 100000000000;

    address owner = msg.sender;
    balances[owner] = totalSupply;

    emit Transfer(0x0, owner, totalSupply);
  }
}
