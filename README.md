# Chainlink hackathon 2024 - GO server
server project used for the chainlink hackathon 2024. this server is used for the cross-chain marketplace.

## Installation
Follow these steps to set up the project on your machine:

1. **Install Go**: 
If you haven't already, [install Go](https://golang.org/dl/) on your machine.

2. **Set up your workspace**: 
Go requires a specific workspace structure. Create a workspace directory and set up the required subdirectories (`src`, `bin`, and `pkg`).

3. **Clone repository**:
git clone https://github.com/cross-chain-market/chainlink-hackathon2024-server.git

4. **Set up environment variables**:
export GO111MODULE=on

5. **Install dependencies**:
cd $HOME/go/src/<repository_name>
go mod tidy

6. **Set up docker and start your docket container**:
make sure docker is installed, and activated.
docker-compose up -d

7. **Run all DB migrations**:
make migrate-drop
make migrate-up

8. **Run server**:
go run ./...


