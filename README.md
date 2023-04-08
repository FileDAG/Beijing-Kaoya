# Beijing-Kaoya


## Discription
Beijing-Kaoya is a Decentralized Storage Network (DSN) based on Ethereum Swarm. We use file incrementation to help save storage costs for multi-version files.

## Environment Requirement

go 1.18

## Install

Be sure to add `$GOPATH/bin` to your terminal's `PATH` if you have not.

Install swarm:
```sh
git clone https://github.com/ethersphere/bee
cd bee
git checkout v1.13.0
make binary
mv dist/bee $GOPATH/bin
```

Install Beijing-Kaoya:
```sh
git clone https://github.com/FileDAG/Beijing-Kaoya.git
cd Beijing-Kaoya
go install
```

To know how to use Beijing-Kaoya just run the following command for help:
```bash
./Beijing-Kaoya
```

