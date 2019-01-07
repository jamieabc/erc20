var AaronToken = artifacts.require("./AaronToken.sol");

module.exports = function(deployer) {
  deployer.deploy(AaronToken);
};
