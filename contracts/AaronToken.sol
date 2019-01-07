pragma solidity 0.4.24;

/* import 'zeppelin-solidity/contracts/token/ERC20/MintableToken.sol'; */
import 'openzeppelin-solidity/contracts/token/ERC20/ERC20.sol';

contract AaronToken is ERC20 {
  string public constant name = "AaronToken";
  string public constant symbol = "AT";
  uint8 public constant decimals = 18;
  uint256 public supplyTokens = 10000;
  uint256 public initialSupply = supplyTokens.mul(10 ** uint256(decimals));
  mapping (address => uint256) public balances;

  constructor() public {
    require(supplyTokens > 0, "amount has to be greater than 0");

    _mint(msg.sender, initialSupply);
    emit Transfer(address(0), msg.sender, initialSupply);
  }
}
