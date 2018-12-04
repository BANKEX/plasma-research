pragma solidity ^0.4.24;

import { RLPReader } from "solidity-rlp/contracts/RLPReader.sol";


library PlasmaDecoder {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  struct Input {
    address owner;
    uint32 blockIndex;
    uint32 txIndex;
    uint8 outputIndex;
    address assetId;
    uint256 amount;
  }

  struct Output {
    address owner;
    address assetId;
    uint256 amount;
  }

  struct Metadata {
    uint32 maxBlockId;
  }

  struct Signature {
    uint256 r;
    uint256 s;
    uint8 v;
  }

  struct Transaction {
    Input[] inputs;
    Output[] outputs;
    Metadata metadata;
    Signature[] signatures;
  }

  struct Block {
    uint32 blockNumber;
    uint256 previousBlockHash;
    uint256 merkleRoot;
    Signature signature;
    Transaction[] transactions;
  }

  function decodeInput(bytes memory rlpBytes) internal pure returns(Input) {
    return _decodeInput(rlpBytes.toRlpItem().toList());
  }

  function decodeOutput(bytes memory rlpBytes) internal pure returns(Output) {
    return _decodeOutput(rlpBytes.toRlpItem().toList());
  }

  function decodeMetadata(bytes memory rlpBytes) internal pure returns(Metadata) {
    return _decodeMetadata(rlpBytes.toRlpItem().toList());
  }

  function decodeSignature(bytes memory rlpBytes) internal pure returns(Signature) {
    return _decodeSignature(rlpBytes.toRlpItem().toList());
  }

  function decodeTransaction(bytes memory rlpBytes) internal pure returns(Transaction) {
    return _decodeTransaction(rlpBytes.toRlpItem().toList());
  }

  function decodeBlock(bytes memory rlpBytes) internal pure returns(Block) {
    return _decodeBlock(rlpBytes.toRlpItem().toList());
  }

  // Private methods

  function _decodeInput(RLPReader.RLPItem[] items) private pure returns(Input) {
    return Input({
      owner: items[0].toAddress(),
      blockIndex: uint32(items[1].toUint()),
      txIndex: uint32(items[2].toUint()),
      outputIndex: uint8(items[3].toUint()),
      assetId: items[4].toAddress(),
      amount: items[5].toUint()
    });
  }

  function _decodeInputs(RLPReader.RLPItem[] memory items) private pure returns(Input[]) {
    Input[] memory inputs = new Input[](items.length);
    for (uint i = 0; i < items.length; i++) {
      inputs[i] = _decodeInput(items[i].toList());
    }
    return inputs;
  }

  function _decodeOutput(RLPReader.RLPItem[] memory items) private pure returns(Output) {
    return Output({
      owner: items[0].toAddress(),
      assetId: items[1].toAddress(),
      amount: items[2].toUint()
    });
  }

  function _decodeOutputs(RLPReader.RLPItem[] memory items) private pure returns(Output[]) {
    Output[] memory outputs = new Output[](items.length);
    for (uint i = 0; i < items.length; i++) {
      outputs[i] = _decodeOutput(items[i].toList());
    }
    return outputs;
  }

  function _decodeMetadata(RLPReader.RLPItem[] memory items) private pure returns(Metadata) {
    return Metadata({
      maxBlockId: uint32(items[0].toUint())
    });
  }

  function _decodeSignature(RLPReader.RLPItem[] memory items) internal pure returns(Signature) {
    return Signature({
      r: items[0].toUint(),
      s: items[0].toUint(),
      v: uint8(items[0].toUint())
    });
  }

  function _decodeSignatures(RLPReader.RLPItem[] memory items) private pure returns(Signature[]) {
    Signature[] memory signatures = new Signature[](items.length);
    for (uint i = 0; i < items.length; i++) {
      signatures[i] = _decodeSignature(items[i].toList());
    }
    return signatures;
  }

  function _decodeTransaction(RLPReader.RLPItem[] memory items) private pure returns(Transaction) {
    return Transaction({
      inputs: _decodeInputs(items[0].toList()),
      outputs: _decodeOutputs(items[1].toList()),
      metadata: _decodeMetadata(items[2].toList()),
      signatures: _decodeSignatures(items[3].toList())
    });
  }

  function _decodeTransactions(RLPReader.RLPItem[] memory items) private pure returns(Transaction[]) {
    Transaction[] memory transactions = new Transaction[](items.length);
    for (uint i = 0; i < items.length; i++) {
      transactions[i] = _decodeTransaction(items[i].toList());
    }
    return transactions;
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