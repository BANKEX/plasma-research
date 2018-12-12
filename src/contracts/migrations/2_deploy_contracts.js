const BankexPlasma = artifacts.require('./BankexPlasma');
const fs = require('fs');
const path = require('path');
const filePath = path.join(__dirname, '../build/BankexPlasma.abi');
const filePathBin = path.join(__dirname, '../build/BankexPlasma.bin');
const BankexPlasmaJSON = path.join(__dirname, '../build/contracts/BankexPlasma.json');

const clean_input = function (str) {
    if ((typeof str === 'number') || (str.isBigNumber === true)) str = `${ str.toString(16) }`
    if ((!str) || (typeof str !== 'string') || (str === '0x')) str = '00'
    if (str.indexOf('0x') === 0) str = str.substr(2)
    if (str.length % 2 === 1) str = `0${str}`
    return `0x${str}`
};

async function WriteBankexPlasmaAbi() {
    const BankexPlasma = JSON.parse(fs.readFileSync(BankexPlasmaJSON, 'utf8'));
    fs.writeFileSync(filePath, JSON.stringify(BankexPlasma.abi));
}

async function WriteBankexPlasmaBin() {
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


