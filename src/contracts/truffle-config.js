const path = require('path');

let mochaConfig = {};
if (process.env.BUILD_TYPE === 'CI') {
  mochaConfig = {
    reporter: 'mocha-junit-reporter',
    reporterOptions: {
      mochaFile: 'result.xml',
    },
  };
}

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  networks: {
    development: {
      host: 'localhost',
      port: 9545,
      network_id: '*', // Match any network id,
      gas: 4600000,
    },
    ganache: {
      host: '127.0.0.1',
      port: 8545,
      network_id: 5777,
      gas: 6721975,
      gasPrice: 1,
    },
    kovan: {
      host: 'localhost',
      port: 8545,
      network_id: 42,
      gas: 4700000,
      gasPrice: 20000000000,
    },
    rinkeby: {
      provider: function () {
        let WalletProvider = require('truffle-wallet-provider');
        let wallet = require('ethereumjs-wallet').fromPrivateKey(Buffer.from(env.ETH_KEY, 'hex'));
        return new WalletProvider(wallet, 'https://rinkeby.infura.io/' + env.INFURA_TOKEN);
      },
      network_id: 4,
    },
    rinkeby_localhost: {
      host: 'localhost', // Connect to geth on the specified
      port: 8545,
      network_id: 4,
      gas: 4612388,
      gasPrice: 20000000000,
      from: '0xf17f52151EbEF6C7334FAD080c5704D77216b732',
    },
    geth_dev: {
      host: 'localhost', // Connect to geth on the specified
      port: 8545,
      network_id: 5777,
      gas: 4700000,
      gasPrice: 20000000000,
    },
  },
  compilers: {
    solc: {
      version: '0.5.2',
      settings: {
        optimizer: {
          enabled: true,
          runs: 200,
        }
      }
    },
  },
  migrations_directory: './migrations',
  contracts_directory: './contracts',
  // eslint-disable-next-line camelcase
  build_directory: path.join(process.cwd(), 'build'),
};
