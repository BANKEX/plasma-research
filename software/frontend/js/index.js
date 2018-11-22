async function deposit(privateKey, amount) {
    const instance = getInstance(ABI, ADDRESS);
    const transactionData = getCallData(instance, "deposit", []);
    const rawTransaction = await signTransaction(privateKey, ADDRESS, amount, transactionData);
    const tx = await sendSignedTransaction(rawTransaction);
    return tx.transactionHash;
}