pragma solidity ^0.4.24;

import "solidity-rlp/contracts/RLPReader.sol";
import { PlasmaTransactions as Tx } from "./PlasmaTransactions.sol";


library PlasmaDecoder {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  struct Block {
    uint32 blockNumber;
    uint256 previousBlockHash;
    uint256 merkleRoot;
    Tx.Signature signature;
    Tx.Transaction[] transactions;
  }

  function rlpItemSize(bytes memory rlpBytes) internal pure returns(uint) {
    return rlpBytes.toRlpItem().toList().length;
  }

  function decodeSegment(bytes memory rlpBytes) internal pure returns(Tx.Input) {
    return _decodeInput(rlpBytes.toRlpItem().toList());
  }

  function decodeInput(bytes memory rlpBytes) internal pure returns(Tx.Input) {
    return _decodeInput(rlpBytes.toRlpItem().toList());
  }

  function decodeOutput(bytes memory rlpBytes) internal pure returns(Tx.Output) {
    return _decodeOutput(rlpBytes.toRlpItem().toList());
  }

  function decodeSignature(bytes memory rlpBytes) internal pure returns(Tx.Signature) {
    return _decodeSignature(rlpBytes.toRlpItem().toList());
  }

  function decodeTransaction(bytes memory rlpBytes) internal pure returns(Tx.Transaction) {
    return _decodeTransaction(rlpBytes.toRlpItem().toList());
  }

  function decodeMerkleProof(bytes memory rlpBytes) internal pure returns(Tx.MerkleProof) {
    return _decodeMerkleProof(rlpBytes.toRlpItem().toList());
  }

  function decodeBlock(bytes memory rlpBytes) internal pure returns(Block) {
    return _decodeBlock(rlpBytes.toRlpItem().toList());
  }

  // Private methods

  function _decodeSegment(RLPReader.RLPItem[] memory items) private pure returns(Tx.Segment) {
    return Tx.Segment({
      begin: uint64(items[0].toUint()),
      end: uint64(items[1].toUint())
    });
  }

  function _decodeInput(RLPReader.RLPItem[] memory items) private pure returns(Tx.Input) {
    return Tx.Input({
      owner: items[0].toAddress(),
      blockIndex: uint64(items[1].toUint()),
      txIndex: uint32(items[2].toUint()),
      outputIndex: uint8(items[3].toUint()),
      amount: _decodeSegment(items[4].toList())
    });
  }

  function _decodeInputs(RLPReader.RLPItem[] memory items) private pure returns(Tx.Input[]) {
    Tx.Input[] memory inputs = new Tx.Input[](items.length);
    for (uint i = 0; i < items.length; i++) {
      inputs[i] = _decodeInput(items[i].toList());
    }
    return inputs;
  }

  function _decodeOutput(RLPReader.RLPItem[] memory items) private pure returns(Tx.Output) {
    return Tx.Output({
      owner: items[0].toAddress(),
      amount: _decodeSegment(items[1].toList())
    });
  }

  function _decodeOutputs(RLPReader.RLPItem[] memory items) private pure returns(Tx.Output[]) {
    Tx.Output[] memory outputs = new Tx.Output[](items.length);
    for (uint i = 0; i < items.length; i++) {
      outputs[i] = _decodeOutput(items[i].toList());
    }
    return outputs;
  }

  function _decodeSignature(RLPReader.RLPItem[] memory items) internal pure returns(Tx.Signature) {
    return Tx.Signature({
      r: items[0].toUint(),
      s: items[1].toUint(),
      v: uint8(items[2].toUint())
    });
  }

  function _decodeSignatures(RLPReader.RLPItem[] memory items) private pure returns(Tx.Signature[]) {
    Tx.Signature[] memory signatures = new Tx.Signature[](items.length);
    for (uint i = 0; i < items.length; i++) {
      signatures[i] = _decodeSignature(items[i].toList());
    }
    return signatures;
  }

  function _decodeTransaction(RLPReader.RLPItem[] memory items) private pure returns(Tx.Transaction) {
    return Tx.Transaction({
      inputs: _decodeInputs(items[0].toList()),
      outputs: _decodeOutputs(items[1].toList()),
      maxBlockIndex: uint64(items[2].toUint()),
      signatures: _decodeSignatures(items[3].toList())
    });
  }

  function _decodeTransactions(RLPReader.RLPItem[] memory items) private pure returns(Tx.Transaction[]) {
    Tx.Transaction[] memory transactions = new Tx.Transaction[](items.length);
    for (uint i = 0; i < items.length; i++) {
      transactions[i] = _decodeTransaction(items[i].toList());
    }
    return transactions;
  }

  function _decodeMerkleProof(RLPReader.RLPItem[] memory items) private pure returns(Tx.MerkleProof) {
    return Tx.MerkleProof({
      proof: items[0].toBytes(),
      root: items[1].toAddress(),
      leaf: items[2].toAddress(),
      index: uint256(items[3].toUint())
    });
  }

  function _decodeBytes32(RLPReader.RLPItem[] memory items) private pure returns(bytes32[]) {
    bytes32[] memory bytesArray = new bytes32[](items.length);
    for (uint i = 0; i < items.length; i++) {
      bytesArray[i] = bytes32(items[i].toUint());
    }
  }

  function _decodeBlock(RLPReader.RLPItem[] memory items) private pure returns(Block) {
    return Block({
      blockNumber: uint32(items[0].toUint()),
      previousBlockHash: items[1].toUint(),
      merkleRoot: items[2].toUint(),
      signature: _decodeSignature(items[3].toList()),
      transactions: _decodeTransactions(items[4].toList())
    });
  }
}
