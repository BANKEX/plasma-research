const { keccak256, bufferToHex } = require('ethereumjs-util');

const empty = keccak256('');

class MerkleTree {
  constructor (elements) {
    // Create layers
    elements = elements.map(el => keccak256(el));
    this.layers = this.getLayers(elements);
  }

  getLayers (elements) {
    let emptyLeveled = empty;
    if (elements.length & 1 == 1) {
      elements.push(emptyLeveled);
    }

    const tree = [elements];
    const maxLevel = 31;
    for (let level = 1; level <= maxLevel; level++) {
      const current = [];
      for (let i = 0; i < tree[level - 1].length / 2; i++) {
        const a = tree[level - 1][i*2];
        const b = tree[level - 1][i*2 + 1];
        const hash = keccak256(Buffer.concat([a, b]));
        current.push(hash);
      }

      if (current.length & 1 && level < maxLevel) {
        current.push(emptyLeveled);
      }
      emptyLeveled = keccak256(Buffer.concat([emptyLeveled, emptyLeveled]));

      tree.push(current);
    }
    return tree;
  }

  getRoot () {
    return this.layers[this.layers.length - 1][0];
  }

  getHexRoot () {
    return bufferToHex(this.getRoot());
  }

  getProof (idx) {
    if (idx === -1) {
      throw new Error('Element does not exist in Merkle tree');
    }
    return this.layers.reduce((proof, layer) => {
      const pairElement = this.getPairElement(idx, layer);
      if (pairElement) {
        proof.push(pairElement);
      }

      idx = Math.floor(idx / 2);
      return proof;
    }, []);
  }

  getHexProof (idx) {
    const result = this.getProof(idx);
    return this.bufArrToHexArr(result);
  }

  getPairElement (idx, layer) {
    const pairIdx = idx % 2 === 0 ? idx + 1 : idx - 1;
    if (pairIdx < layer.length) {
      return layer[pairIdx];
    } else {
      return null;
    }
  }

  bufIndexOf (el, arr) {
    let hash;

    // Convert element to 32 byte hash if it is not one already
    if (el.length !== 32 || !Buffer.isBuffer(el)) {
      hash = keccak256(el);
    } else {
      hash = el;
    }

    for (let i = 0; i < arr.length; i++) {
      if (hash.equals(arr[i])) {
        return i;
      }
    }

    return -1;
  }

  bufArrToHexArr (arr) {
    if (arr.some(el => !Buffer.isBuffer(el))) {
      throw new Error('Array is not an array of buffers');
    }

    return arr.map(el => '0x' + el.toString('hex'));
  }
}

module.exports = {
  MerkleTree,
};
