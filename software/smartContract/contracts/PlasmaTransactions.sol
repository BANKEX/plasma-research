pragma solidity ^0.4.24;

import "./PlasmaDecoder.sol";
import { MerkleProof as MerkleProofLib } from "./MerkleProof.sol";


library PlasmaTransactions {

  struct Segment {
    uint64 begin;
    uint64 end;
  }

  struct Input {
    address owner;
    uint64 blockIndex;
    uint32 txIndex;
    uint8 outputIndex;
    Segment amount;
  }

  struct Output {
    address owner;
    Segment amount;
  }

  struct Signature {
    uint256 r;
    uint256 s;
    uint8 v;
  }

  struct Transaction {
    Input[] inputs;
    Output[] outputs;
    uint64 maxBlockIndex;
    Signature[] signatures;
  }

  struct MerkleProof {
    bytes32[] proof;
    bytes32 root;
    bytes32 leaf;
    uint256 index;
  }

  function validateTransaction(bytes memory rlpTransaction) internal pure returns (bool) {
    Transaction memory tr = PlasmaDecoder.decodeTransaction(rlpTransaction);

    require(tr.inputs.length != 0, "Empty inputs");
    require(tr.outputs.length != 0, "Empty outputs");
    require(tr.inputs.length <= tr.outputs.length, "Wrong outputs size");

    for (uint i = 0; i < tr.inputs.length; i += 1) {
      Input memory input = tr.inputs[i];
      _validateInput(input);
    }

    for (i = 0; i < tr.outputs.length; i += 1) {
      Output memory output = tr.outputs[i];
      _validateOutput(output);
    }

    require(tr.maxBlockIndex > 0, "Wrong max block index");

    require(tr.signatures.length > 1, "Wrong signatures count");
    for (i = 0; i < tr.signatures.length; i += 1) {
      Signature memory sign = tr.signatures[i];
      require(sign.r > 0, "Wrong `r` at signature");
      require(sign.s > 0, "Wrong `s` at signature");
      require(sign.v > 0, "Wrong `v` at signature");
    }

    return true;
  }

  function validateInputOutput(bytes memory rlpInputOutput, bytes memory rlpMerkleProof) internal pure returns (bool) {
    uint size = PlasmaDecoder.rlpItemSize(rlpInputOutput);
    require(size == 5 || size == 2, "Wrong input/output size");

    Input memory input;
    Output memory output;
    if (size == 5) {
      input = PlasmaDecoder.decodeInput(rlpInputOutput);
      _validateInput(input);
    } else if (size == 2) {
      output = PlasmaDecoder.decodeOutput(rlpInputOutput);
      _validateOutput(output);
    }

    MerkleProof memory proof = PlasmaDecoder.decodeMerkleProof(rlpMerkleProof);
    return _verifyMerkleProof(proof);
  }

  function _verifyMerkleProof(MerkleProof p) private pure returns(bool) {
    return MerkleProofLib.verifyAtIndex(p.proof, p.root, p.leaf, p.index);
  }

  function _validateInput(Input memory input) private pure returns(bool) {
    require(input.owner != address(0), "Empty owner address");
    require(input.blockIndex >= 0, "Wrong block index at input");
    require(input.txIndex >= 0, "Wrong transaction index at input");
    require(input.outputIndex >= 0, "Wrong output index at input");
    Segment memory amount = input.amount;
    require(amount.begin >= 0, "Wrong segment amount begin number at input");
    require(amount.end >= 0, "Wrong segment amount end number at input");
    require(amount.begin != amount.end, "Segment begin equals to end number at input");
    return true;
  }

  function _validateOutput(Output memory output) private pure returns(bool) {
    require(output.owner != address(0), "Empty owner address at output");
    Segment memory amount = output.amount;
    require(amount.begin >= 0, "Wrong segment amount begin number at output");
    require(amount.end >= 0, "Wrong segment amount end number at output");
    require(amount.begin != amount.end, "Segment begin equals to end number at output");
    return true;
  }

}
