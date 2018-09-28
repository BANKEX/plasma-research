# plasma-research
That repo contains docs and code snippets related to BANKEX Multi-Asset Plasma implementation.

# Multiasset plasma

The implementation that we are going to make is based on [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) proposal. 
Additionally supports it will support deposit, withdraw, transfer and exchange of a wide range of Ethereum tokens standards like ERC20, ERC721, ERC888, etc..
And that is what we actually mean by Multi-asset.

To support that we are going to implement the more sophisticated parent smart contract that will be able to take ownership of different tokens types.
Although, Plasma owner can't guarantee that actual token contract implemented safely and fully correctly. 
For the end user, it means that before the trade and safe
That means that is the responsibility of the user to ensure that actual asset he going to trade is safe.

All the research we do, we going to document in the form of Yellow Paper.
At the moment we have following separate docs that cover different parts of our application:
- [Architecture overview](https://github.com/BANKEX/plasma-research/blob/master/docs/architecture.md)
- [Block structure](https://github.com/BANKEX/plasma-research/blob/master/docs/block-structure.md)

