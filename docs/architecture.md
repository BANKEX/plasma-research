# Open Plasma Architecture

## Description

## Plasma owner application

### Block structure

## Plasma client application

## Ethereum Smart Contract

## Appendixes

### Exit game

We suppose to use standard [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) exit game with some additions. As we are going to support atomic swaps with multiple transaction source owners, one of a participant may have not all signatures of the transaction. So, we need to set up one more exit game branch with signatures collection.

![exit game schema](https://raw.githubusercontent.com/BANKEX/plasma-research/master/docs/assets/plasma_exit_game.svg)
