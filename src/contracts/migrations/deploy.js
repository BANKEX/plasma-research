const fs = require('fs');
const path = require('path');
// const filePath = path.join(__dirname, '../../nodde/config/config.verifier.json');
const BankexPlasmaJSON = path.join(__dirname, '../build/contracts/BankexPlasma.json');
const Tx = require('ethereumjs-tx');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));

const clean_input = function (str) {
    if ((typeof str === 'number') || (str.isBigNumber === true)) str = `${ str.toString(16) }`
    if ((!str) || (typeof str !== 'string') || (str === '0x')) str = '00'
    if (str.indexOf('0x') === 0) str = str.substr(2)
    if (str.length % 2 === 1) str = `0${str}`
    return `0x${str}`
};

(async function deploy() {
    const BankexPlasma = JSON.parse(fs.readFileSync(BankexPlasmaJSON, 'utf8'));
    const bytecode = BankexPlasma.bytecode;
    // const data = JSON.parse(fs.readFileSync(filePath, 'utf8'));
    const deployerAddress = "0x9cA4E1F69A3ABD60989864FAd1025095dFCC58F1";
    const privateKey = Buffer.from("e4058d9c3a81b4e95d8e3a17a5f52486a7fc411e57dcd4f6c771dbc2428928e9", "hex");
    const txParams = {
        gasPrice: clean_input(1),
        gas: clean_input(6050992),
        from: clean_input(deployerAddress),
        nonce: clean_input(await web3.eth.getTransactionCount(deployerAddress)),
        value: clean_input(0),
        data: clean_input(bytecode)
    };
    const tx = new Tx(txParams);
    tx.sign(privateKey);
    const serializedTx = tx.serialize();
    const txHash = await web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'));
    console.log("BankexPlasma address: " + txHash.contractAddress);
    // data.plasma_contract_address = txHash.contractAddress;
    // fs.writeFileSync(filePath, JSON.stringify(data));
})();