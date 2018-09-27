const PlasmaMain = artifacts.require('PlasmaMain');

module.exports = async function (deployer) {
    deployer.deploy(PlasmaMain);
};
