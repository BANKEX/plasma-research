const rlp = require('rlp');
const { bufferToHex } = require('ethereumjs-util');
const { MerkleTree, keccak160 } = require('./helpers/merkleTree.js');
const { expectThrow } = require('./helpers/expectThrow');

const PlasmaTransactionsWrapper = artifacts.require('PlasmaTransactionsWrapper');
const MerkleProofWrapper = artifacts.require('MerkleProofWrapper');

require('chai')
  .should();

contract('PlasmaTransactions', (accounts) => {
  let toHex = (buf) => {
    buf = buf.toString('hex');
    if (buf.substring(0, 2) == "0x")
        return buf;
    return "0x" + buf.toString("hex");
  };

  let plasmaTransactions;
  let merkleProofWrapper;


  before('before', async () => {
    plasmaTransactions = await PlasmaTransactionsWrapper.deployed();
    merkleProofWrapper = await MerkleProofWrapper.deployed();
  });

  describe('validateTransaction', () => {
    it('should return `true` for valid transaction', async () => {
      let tx = [
        [
          [accounts[0], 1, 1, 0, [2, 3]],
          [accounts[1], 1, 1, 1, [2, 3]]
        ],
        [
          [accounts[0], [2, 3]],
          [accounts[1], [2, 3]]
        ],
        1,
        [
          [1, 2, 3],
          [4, 5, 6]
        ]
      ];
      tx = toHex(rlp.encode(tx));
      (await plasmaTransactions.validateTransaction(tx)).should.equal(true);
    });

    it('should throw exception for invalid transaction with inputs > outputs', async () => {
      let tx = [
        [
          [accounts[0], 1, 1, 0, [2, 3]],
          [accounts[1], 1, 1, 1, [2, 3]],
          [accounts[1], 1, 1, 1, [2, 3]]
        ],
        [
          [accounts[0], [2, 3]],
          [accounts[1], [2, 3]]
        ],
        1,
        [
          [1, 2, 3],
          [4, 5, 6]
        ]
      ];
      tx = toHex(rlp.encode(tx));
      await expectThrow(plasmaTransactions.validateTransaction(tx));
    });
  });

  describe('validateInputOutput', () => {
    it('should return `true` for valid input', async () => {
      let input = [accounts[0], 1, 1, 0, [2, 3]];
      input = toHex(rlp.encode(input));
      (await plasmaTransactions.validateInputOutput(input)).should.equal(true);
    });
    it('should throw exception for invalid input format', async () => {
      let input = [0, 1, 0, [2, 3]];
      input = toHex(rlp.encode(input));
      await expectThrow(plasmaTransactions.validateInputOutput(input));
    });
    it('should return `true` for valid output', async () => {
      let output = [accounts[1], [2, 3]];
      output = toHex(rlp.encode(output));
      (await plasmaTransactions.validateInputOutput(output)).should.equal(true);
    });
    it('should throw exception for invalid output format', async () => {
      let output = [0, 1, [2, 3]];
      output = toHex(rlp.encode(output));
      await expectThrow(plasmaTransactions.validateInputOutput(output));
    });
  });

  describe('verifyMerkleProof', () => {
    it('should return `true` for valid MerkleProof', async () => {
      const elements = [];
      const size = 6;
      for (let i = 0; i < size; i += 1) {
        elements.push(`a${i}`);
      }
      const merkleTree = new MerkleTree(elements, keccak160);
      const root = merkleTree.getHexRoot();
      for (let i = 0; i < size; i += 1) {
        var proof = merkleTree.getProof(i);
        proof = bufferToHex(Buffer.concat(proof));
        const leaf = bufferToHex(merkleTree.layers[0][i]);
        let merkleProof = [proof, root, leaf, i];
        merkleProof = toHex(rlp.encode(merkleProof).toString('hex'));
        (await plasmaTransactions.verifyMerkleProof(merkleProof)).should.equal(true);
      }
    });
    it('should return `false` for invalid MerkleProof', async () => {
      const correctElements = ['a', 'b', 'c'];
      const correctMerkleTree = new MerkleTree(correctElements, keccak160);
      const correctRoot = correctMerkleTree.getHexRoot();
      const correctLeaf = bufferToHex(keccak160(correctElements[0]));

      const badElements = ['d', 'e', 'f', ''];
      const badMerkleTree = new MerkleTree(badElements);
      const badProof = badMerkleTree.getHexProof(badElements[0]);

      let merkleProof = [badProof, correctRoot, correctLeaf, 0];
      merkleProof = toHex(rlp.encode(merkleProof));
      (await plasmaTransactions.verifyMerkleProof(merkleProof)).should.equal(false);
    });
  });
});
