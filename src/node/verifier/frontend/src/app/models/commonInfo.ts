export interface CommonInfo {
  contractAddress: string;
  contractBalance: string;
  verifierEtherBalance: string;
  verifierPlasmaBalance: string;
  latestBlock: string;
  verifierInputs: Input[];
}

export interface Input {
  blockNumber: any;
  txNumber: any;
  outputNumber: any;
  output: Output;
}

export interface Output {
  owner: any;
  slice: Slice;
}

export interface Slice {
  begin: any;
  end: any;
}
