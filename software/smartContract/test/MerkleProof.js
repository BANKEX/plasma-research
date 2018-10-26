const { MerkleTree } = require('./helpers/merkleTree.js');
const { keccak256, bufferToHex } = require('ethereumjs-util');

const MerkleProofWrapper = artifacts.require('MerkleProofWrapper');

require('chai')
    .should();

function ds (size) {
    it(`should return true for a valid Merkle proof with ${size} elements`, async function () {
        const elements = [];
        for (let i = 0; i < size; i += 1) {
            elements.push(`a${i}`);
        }
        const merkleTree = new MerkleTree(elements);
        const root = merkleTree.getHexRoot();
        for (let i = 0; i < size; i += 1) {
            const proof = merkleTree.getHexProof(i);
            const leaf = bufferToHex(merkleTree.layers[0][i]);
            (await this.merkleProof.verifyAtIndex(proof, root, leaf, i)).should.equal(true);
        }
    });
}

contract('MerkleProof', function () {
    beforeEach(async function () {
        this.merkleProof = await MerkleProofWrapper.new();
    });

    describe('verifyAtIndex', function () {
        for (let i = 1; i < 32; i += 1) {
            ds(i);
        }

        it('should return false for an invalid Merkle proof', async function () {
            const correctElements = ['a', 'b', 'c'];
            const correctMerkleTree = new MerkleTree(correctElements);

            const correctRoot = correctMerkleTree.getHexRoot();

            const correctLeaf = bufferToHex(keccak256(correctElements[0]));

            const badElements = ['d', 'e', 'f', ''];
            const badMerkleTree = new MerkleTree(badElements);

            const badProof = badMerkleTree.getHexProof(badElements[0]);

            (await this.merkleProof.verifyAtIndex(badProof, correctRoot, correctLeaf, 0)).should.equal(false);
        });

        it('should return false for a Merkle proof of invalid length', async function () {
            const elements = ['a', 'b', 'c'];
            const merkleTree = new MerkleTree(elements);

            const root = merkleTree.getHexRoot();

            const proof = merkleTree.getHexProof(elements[0]);
            const badProof = proof.slice(0, proof.length - 5);

            const leaf = bufferToHex(keccak256(elements[0]));

            (await this.merkleProof.verifyAtIndex(badProof, root, leaf, 0)).should.equal(false);
        });
    });
});
