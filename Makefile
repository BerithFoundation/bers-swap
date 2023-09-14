build:
	go build -o run_test/berith-swap

runtest:
	cd run_test && ./berith-swap

solc:
	cd contract && npx hardhat compile

deploy:
	cd contract && npx hardhat run scripts/bers-token.js --network klaytnTestnet
	cd contract && npx hardhat run scripts/berith-swap.js --network berithTestnet

ctest:
	cd contract && npx hardhat test

test:
	go test ./... -timeout=3m -v -short -cover

sqlc:
	sqlc generate

.PHONY: build runtest solc deploy ctest test sqlc
