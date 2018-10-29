const { MerkleTree, keccak160 } = require('./helpers/merkleTree.js');
const { keccak256, bufferToHex } = require('ethereumjs-util');

const MerkleProofWrapper = artifacts.require('MerkleProofWrapper');

require('chai')
    .should();

function verifyTree (size) {
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

function verifyTree160 (size) {
    it(`should return true for a valid Merkle proof with ${size} elements`, async function () {
        const elements = [];
        for (let i = 0; i < size; i += 1) {
            elements.push(`a${i}`);
        }
        const merkleTree = new MerkleTree(elements, keccak160);
        const root = merkleTree.getHexRoot();
        for (let i = 0; i < size; i += 1) {
            var proof = merkleTree.getProof(i);
            proof = bufferToHex(Buffer.concat(proof));
            const leaf = bufferToHex(merkleTree.layers[0][i]);
            (await this.merkleProof.verifyAtIndex160(proof, root, leaf, i)).should.equal(true);
        }
    });
}

contract('MerkleProof', function () {
    beforeEach(async function () {
        this.merkleProof = await MerkleProofWrapper.new();
    });

    describe('verifyAtIndex160', function () {
        // verifyTree160(4);
        for (let size = 1; size < 6; size += 1) {
            verifyTree160(size);
        }

        it('should return false for an invalid Merkle proof', async function () {
            const correctElements = ['a', 'b', 'c'];
            const correctMerkleTree = new MerkleTree160(correctElements);

            const correctRoot = correctMerkleTree.getHexRoot();

            const correctLeaf = bufferToHex(keccak160(correctElements[0]));

            const badElements = ['d', 'e', 'f', ''];
            const badMerkleTree = new MerkleTree(badElements);

            const badProof = badMerkleTree.getHexProof(badElements[0]);

            (await this.merkleProof.verifyAtIndex160(badProof, correctRoot, correctLeaf, 0)).should.equal(false);
        });

        it('should return false for a Merkle proof of invalid length', async function () {
            const elements = ['a', 'b', 'c'];
            const merkleTree = new MerkleTree160(elements);

            const root = merkleTree.getHexRoot();

            const proof = merkleTree.getHexProof(elements[0]);
            const badProof = proof.slice(0, proof.length - 5);

            const leaf = bufferToHex(keccak160(elements[0]));

            (await this.merkleProof.verifyAtIndex160(badProof, root, leaf, 0)).should.equal(false);
        });
    });

    describe('verifyAtIndex', function () {
        for (let size = 1; size < 6; size += 1) {
            verifyTree(size);
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
