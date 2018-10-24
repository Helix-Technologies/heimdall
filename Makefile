dep:
	dep ensure -v
	mkdir -p vendor/github.com/tendermint
	git clone -b v0.9.2 --single-branch --depth 1 https://github.com/tendermint/iavl vendor/github.com/tendermint/iavl

clean:
	rm -rf build

build: clean
	mkdir -p build
	rm -rf vendor/github.com/ethereum/go-ethereum
	git clone -b v1.8.17 --single-branch --depth 1 https://github.com/ethereum/go-ethereum vendor/github.com/ethereum/go-ethereum
	go build -o build/heimdalld cmd/heimdalld/main.go
	go build -o build/heimdallcli cmd/heimdallcli/main.go

contracts:
	# mkdir -p contracts/validatorset
	# abigen --abi=contracts/validatorset/validatorset.abi --pkg=validatorset --out=contracts/validatorset/ValidatorSet.go

run-heimdall:
	./build/heimdalld start

rest-server:
	./build/heimdallcli rest-server

start: build run-heimdall

#
# docker commands
#

build-docker:
	cd docker; make build

build-docker-develop:
	cd docker; make build-develop

.PHONY: dep build
