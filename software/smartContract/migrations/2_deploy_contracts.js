const PlasmaBlocks = artifacts.require('PlasmaBlocks');

module.exports = async function (deployer) {
    deployer.deploy(PlasmaBlocks);
};
