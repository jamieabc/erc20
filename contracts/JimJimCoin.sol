pragma solidity ^0.4.24;

import 'zeppelin-solidity/contracts/token/ERC20/MintableToken.sol';

contract JimJimCoin is MintableToken {

  string public name = "JimJimCoin";
  string public symbol = "JJC";
  uint public decimals = 18;
  uint public INITIAL_SUPPLY = 10000 * (10 ** decimals);

  constructor() public {
    totalSupply_ = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
    emit Transfer(address(0x0), msg.sender, INITIAL_SUPPLY);
  }
}
