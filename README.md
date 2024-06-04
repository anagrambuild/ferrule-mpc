# Ferrule MPC

## WIP - Work in Progress
Not recommended for production use. This project is still in development and may not be stable or fully functional.

## Introduction

Ferrule MPC is a golang library/tool to run a Farcaster Wallet App or Connected App as an MPC clustered signer.
The cluster is leader less and uses a threshold signature scheme to sign transactions with a threshold of `t` out of `n` signers. It uses the binance tss library for the threshold signature scheme and the memberlist library for the cluster membership which uses the SWIM protocol. 

## Goals

- Make it eaiser to deploy decentralized Farcaster apps
- Be the foundational signer technology for Ferrule (the Farcaster Hub AVS)
- Leaderless consensus

## Getting Started

Currently you can run this by using docker compose. For a local testnet all you need is to run the following command:

```bash
docker-compose up
```

This will start a 3 node cluster with a threshold of 2 out of 3 signers. You can see all environment vars in the `docker-compose.yml` file.

## Particulars 

Currently the cluster uses a basic consensus system that assumes all signers are honest. This is not secure in a production environment. We are working on a more secure consensus system that will be implemented in the future. 
