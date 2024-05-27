# Cross-Chain Marketplace

This repository contains the backend server for the Cross-Chain Marketplace. The server handles various marketplace functionalities, processes smart contract events, and provides user authentication features.

> _Server project used for the ChainLink Block Magic Hackathon_

## Features

- **Marketplace Management**: Manage collections, items, list/unlist items, buy/sell items.
- **Smart Contract Event Listener**: Listen for and process updates from smart contract events.
- **Locking System**: Prevent two buyers from purchasing the same item simultaneously.
- **User Authentication**: Simple user login and registration.

## Technologies Used

- **Golang 1.22**
- **Geth** library (https://github.com/ethereum/go-ethereum)
- **PostgreSQL**
- **Docker & Docker Compose**

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Geth](https://github.com/ethereum/go-ethereum)

## Installation
Follow these steps to set up the project on your machine:

1. **Install Go**: 
If you haven't already, [install Go](https://golang.org/dl/) on your machine.

2. **Clone repository**:
```sh
git clone https://github.com/cross-chain-market/chainlink-hackathon2024-server.git
```

3. **Install dependencies**:
 
```sh
cd chainlink-hackathon2024-server
go mod tidy
```

4. **Set up docker and start your docket container**:

```sh
make sure docker is installed, and activated.
docker-compose up -d
```

5. **Run all DB migrations**:

```sh
make migrate-drop
make migrate-up
```

6. **Run server**:

```sh
go run ./...
```

## How it works?

### Endpoints

The backend server exposes several endpoints to interact with the marketplace and user authentication features. Below are some key endpoints:

- **Marketplace Endpoints**
    - `POST /v1/users/{userAddress}/collections`: Create a new Collection.
    - `GET /v1/users/{userAddress}/collections`: Get User's Collections.
    - `GET /v1/users/{userAddress}/collections/{collectionId}`: Get Collection by ID.
    - `POST /v1/users/{userAddress}/collections/{collectionId}/items/{itemId}/list`: List an Item for sale.
    - `POST /v1/users/{userAddress}/collections/{collectionId}/items/{itemId}/unlist`: Unlist an Item.
    - `GET /v1/listings?collectionId={collectionId}`: Get active Listings.
    - `POST /v1/users/{userAddress}/collections/{collectionId}/items/{itemId}/buy`: Buy an Item.

- **User Authentication Endpoints**
    - `POST /v1/users/register`: Register a new User.
    - `POST /v1/users/login`: Login an existing User.
    - `GET /v1/users/{userEmail}`: Get User information.

### Event Listener

The server listens for smart contract events and processes updates accordingly. This ensures the marketplace reflects the current state of the blockchain.

### Locking System

A robust locking system is implemented to avoid race conditions where two buyers might try to purchase the same item simultaneously.