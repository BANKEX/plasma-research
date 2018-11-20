const PlasmaBlocks = artifacts.require('PlasmaBlocks');
const MerkleProof = artifacts.require('MerkleProof');
const MerkleProofWrapper = artifacts.require('MerkleProofWrapper');
const PlasmaDecoder = artifacts.require('PlasmaDecoder');
const PlasmaTransactions = artifacts.require('PlasmaTransactions');
const PlasmaTransactionsWrapper = artifacts.require('PlasmaTransactionsWrapper');

module.exports = deployer => {
  deployer.then(async () => {
    await deployer.deploy(PlasmaBlocks);
    await deployer.deploy(MerkleProof);
    await deployer.deploy(MerkleProofWrapper);
    await deployer.deploy(PlasmaDecoder);
    await deployer.deploy(PlasmaTransactions);
    await deployer.deploy(PlasmaTransactionsWrapper);
  });
}
