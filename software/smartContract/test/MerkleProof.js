const { MerkleTree } = require('./helpers/merkleTree.js');
const { sha3, bufferToHex } = require('ethereumjs-util');

const MerkleProofWrapper = artifacts.require('MerkleProofWrapper');

require('chai')
.should();

function ds(size) {
  it(`should return true for a valid Merkle proof with ${size} elements`, async function () {
    var elements = []
    for (var i = 0; i < size; i += 1) {
      elements.push(`a${i}`)
    }
    while ((elements.length & (elements.length - 1)) != 0) {
      elements.push('');
    }
    const merkleTree = new MerkleTree(elements);
    const root = merkleTree.getHexRoot();

    for (var i = 0; i < size; i += 1) {
      const proof = merkleTree.getHexProof(elements[i]);
      const leaf = bufferToHex(merkleTree.elements[i]);
      (await this.merkleProof.verify(proof, root, leaf, i)).should.equal(true);
    }
  });
}

contract('MerkleProof', function () {
  beforeEach(async function () {
    this.merkleProof = await MerkleProofWrapper.new();
  });

  describe('verify', function () {
    for (var i = 1; i < 32; i += 1) {
      ds(i)
    }

    it('should return false for an invalid Merkle proof', async function () {
      const correctElements = ['a', 'b', 'c', ''];
      const correctMerkleTree = new MerkleTree(correctElements);

      const correctRoot = correctMerkleTree.getHexRoot();

      const correctLeaf = bufferToHex(sha3(correctElements[0]));

      const badElements = ['d', 'e', 'f', ''];
      const badMerkleTree = new MerkleTree(badElements);

      const badProof = badMerkleTree.getHexProof(badElements[0]);

      (await this.merkleProof.verify(badProof, correctRoot, correctLeaf, 0)).should.equal(false);
    });

    it('should return false for a Merkle proof of invalid length', async function () {
      const elements = ['a', 'b', 'c', ''];
      const merkleTree = new MerkleTree(elements);

      const root = merkleTree.getHexRoot();

      const proof = merkleTree.getHexProof(elements[0]);
      const badProof = proof.slice(0, proof.length - 5);

      const leaf = bufferToHex(sha3(elements[0]));

      (await this.merkleProof.verify(badProof, root, leaf, 0)).should.equal(false);
    });
  });
});
