const BankexPlasma = artifacts.require('BankexPlasma');

module.exports = async function (deployer) {
  deployer.deploy(BankexPlasma);
};
