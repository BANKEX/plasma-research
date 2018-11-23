async function deposit(privateKey, amount) {
    const instance = getInstance(ABI, ADDRESS);
    const transactionData = getCallData(instance, "deposit", []);
    const rawTransaction = await signTransaction(privateKey, ADDRESS, amount, transactionData);
    const tx = await sendSignedTransaction(rawTransaction);
    return tx.transactionHash;
}

async function depositViaMetaMask(amount) {
    const instance = getInstance(ABI, ADDRESS);
    const transactionData = getCallData(instance, "deposit", []);
    const tx = await sendTransactionViaMetaMask(ADDRESS, amount, transactionData);
    return tx.transactionHash;
}

