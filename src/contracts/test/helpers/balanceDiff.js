async function balanceDifference (account, promise) {
  const balanceBefore = web3.eth.getBalance(account);
  await promise();
  const balanceAfter = web3.eth.getBalance(account);
  return balanceAfter.minus(balanceBefore);
}

async function balanceDifferenceRespectingFees (target, options) {
  const balanceBefore = await web3.eth.getBalance(options.from);
  const { receipt } = await target.sendTransaction(options);
  const balanceAfter = await web3.eth.getBalance(options.from);

  const gasPrice = options.gasPrice || receipt.gasPrice || web3.eth.gasPrice;
  const fees = (new web3.BigNumber(receipt.gasUsed)).mul(new web3.BigNumber(gasPrice));
  return balanceAfter.add(options.value).sub(balanceBefore.sub(fees));
}

module.exports = {
  balanceDifference,
  balanceDifferenceRespectingFees,
};
