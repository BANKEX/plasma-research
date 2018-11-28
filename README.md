# plasma-research
That repo contains docs and code snippets related to BANKEX Multi-Asset Plasma implementation.

# Multiasset plasma

The implementation that we are going to make is based on [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) proposal. 
Additionally supports it will support deposit, withdraw, transfer and exchange of a wide range of Ethereum tokens standards like ERC20, ERC721, ERC888, etc..
And that is what we actually mean by Multi-asset.

To support that we are going to implement the more sophisticated parent smart contract that will be able to take ownership of different tokens types.
Although, Plasma owner can't guarantee that actual token contract is implemented safely and fully correctly.
For the end user that means that he or she should check that token he was offered doesn't have a bugs and fraud in implementation by himself.

Each token type that was deposit to plasma smart contract will get a unique id. That id will be used further in the transactions to describe transfer and exchange operations.
To implement an exchange operation, we need a single transaction signed by both parties. The protocol that will allow users to share transactions for signatures will be mainly implemented in the plasma client.

All the research we do, we going to document in the form of Yellow Paper.
At the moment we have following separate docs that cover different parts of our application:
- [Architecture overview](https://github.com/BANKEX/plasma-research/blob/master/docs/architecture.md)
- [Block structure](https://github.com/BANKEX/plasma-research/blob/master/docs/block-structure.md)

# Installing

1. go to /software directory
2. run install.sh
3. run run.sh
4. open remix
5. connect to web3 provider http://localhost:9545
6. copy this code:
``` js
 pragma solidity 0.4.25;
 contract BankexPlasma {
 function deposit() public payable {}
 }
```
7. Compile it and go ti Run tab
8. Copy Plasma Smart contract address from console where you run your 'run.sh' file or you can write 'smartContractAddress' at CLI
8. Put your this addresss to input 'At Address' and push the button
9. Push deposit button nearby
10. Whatch console where your run 'run.sh'. You can write events to see all deposit events

