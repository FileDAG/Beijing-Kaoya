# Beijing Kaoya


## Discribe
This software is based on ethereum swarm, a DSN. We apply file increment on it to help save 
storage cost for multi-version files. 

## Environment Requirement

go 1.18

## Install
To use Beijing Kaoya, you need to install swarm first:
```sh
git clone https://github.com/ethersphere/bee
cd bee
git checkout v1.13.0
make binary
```
swarm client will be installed in directory dist.

Next, you can install Beijing Kaoya:
```sh
git clone https://github.com/FileDAG/Beijing-Kaoya.git
cd Beijing-Kaoya
go build
```
We recommand you to put Beijing-Kaoya, dev_start.sh and swarm client into a same directory so 
you can use Beijing-Kaoya to start and fund a swarm node.

To know how tu use Beijing-Kaoya just run the following command for help:
```bash
./Beijing-Kaoya
```

