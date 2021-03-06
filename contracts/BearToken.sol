pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

/**
 * @title BeatToken is a basic ERC20 Token
 */
contract BearToken is ERC20 {
  uint256 public totalSupply;
  string public name;
  string public symbol;
  uint32 public decimals;
  mapping(address => uint256) public balances;

  constructor() public {
    symbol = "BEAR";
    name = "BearToken";
    decimals = 5;
    totalSupply = 100000000000;

    address owner = msg.sender;
    balances[owner] = totalSupply;

    emit Transfer(0x0, owner, totalSupply);
  }
}
