const BankexPlasma = artifacts.require('./BankexPlasma');
const fs = require('fs');
const path = require('path');
const filePath = path.join(__dirname, '../build/BankexPlasma.abi');
const filePathBin = path.join(__dirname, '../build/BankexPlasma.bin');
const BankexPlasmaJSON = path.join(__dirname, '../build/contracts/BankexPlasma.json');

async function WriteBankexPlasmaAbi () {
  const BankexPlasma = JSON.parse(fs.readFileSync(BankexPlasmaJSON, 'utf8'));
  fs.writeFileSync(filePath, JSON.stringify(BankexPlasma.abi));
}

async function WriteBankexPlasmaBin () {
  const BankexPlasma = JSON.parse(fs.readFileSync(BankexPlasmaJSON, 'utf8'));
  fs.writeFileSync(filePathBin, BankexPlasma.bytecode);
}

module.exports = function (deployer) {
  deployer.then(async function () {
    const contract = await BankexPlasma.new();
    WriteBankexPlasmaAbi();
    WriteBankexPlasmaBin();
    console.log('BankexPlasma address: ' + contract.address);
  }).catch(function () {
  });
};
