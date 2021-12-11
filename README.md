[![Go](https://github.com/oluwakeye-john/wallet-alert/actions/workflows/go.yml/badge.svg)](https://github.com/oluwakeye-john/wallet-alert/actions/workflows/go.yml)

<img src="static/images/bell.svg" width="150px" />

# Wallet Alert

Get notified(via email) when a transaction occurs in your wallet address.

## Supported Currencies

- Bitcoin
- Ethereum
- Dogecoin
- Dash
- Litecoin
- Blockcypher(Test)

## Integrations

- BlockCypher
- Postgres
- Mailtrap

## Screenshots

![image](https://user-images.githubusercontent.com/43508135/145659869-d175992c-569d-45d0-8668-b73def058ed5.png)


## Online Demo

https://wallet-alert.herokuapp.com

## Running locally

#### Clone the repo and download dependencies

```
git clone https://github.com/oluwakeye-john/wallet-alert.git
cd wallet-alert
go mod tidy
```

#### Create a `.env` file and populate it, using the `.env.example` file as a reference

```
touch .env
```

#### Start the project

```
make dev
```

## Running tests

```
make test
```
