const tbn = (x) => new BigNumber(x);
const tw = (x) => BigNumber.isBigNumber(x) ? x.times(1e18).integerValue() : tbn(x).times(1e18).integerValue();
const fw = (x) => BigNumber.isBigNumber(x) ? x.times(1e-18).toNumber() : tbn(x).times(1e-18).toNumber();

if (typeof window.web3 !== 'undefined') {
    // Use Mist/MetaMask's provider
    window.web3 = new Web3(web3.currentProvider);
} else {
    window.web3 =  new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:9545'));
}


async function get(instance, method, parameters) {
    return await instance.methods[method](...parameters).call();
}

async function signTransaction(privateKey, receiver, amount, transactionData) {
    const userAddress = getAddress(privateKey);
    const txParam = {
        nonce: Number(await web3.eth.getTransactionCount(userAddress)),
        to: receiver,
        value: Number(amount),
        from: userAddress,
        data: transactionData !== undefined ? transactionData : '',
        gasPrice: await web3.eth.getGasPrice(),
        gas: 210000
    };
    const privateKeyBuffer = ethereumjs.Buffer.Buffer.from(privateKey.substring(2), 'hex');

    const tx = new ethereumjs.Tx(txParam);
    tx.sign(privateKeyBuffer);
    const serializedTx = tx.serialize();
    return '0x' + serializedTx.toString('hex');
}

async function sendTransactionViaMetaMask(receiver, amount, transactionData) {
    const txParam = {
        to: receiver,
        from: web3.eth.accounts.givenProvider.selectedAddress,
        value: amount,
        data: transactionData !== undefined ? transactionData : '',
        gasPrice: await web3.eth.getGasPrice(),
        gas: 210000
    };

    return await web3.eth.sendTransaction(txParam)
}

async function sendSignedTransaction(rawTransaction) {
    return await web3.eth.sendSignedTransaction(rawTransaction);
}

function getCallData(instance, method, parameters) {
    return instance.methods[method](...parameters).encodeABI();
}

function getInstance(ABI, address) {
    return new web3.eth.Contract(ABI, address);
}

function getAddress(privateKey) {
    let _privateKey = privateKey.substring(2, privateKey.length);
    return keythereum.privateKeyToAddress(_privateKey);
}