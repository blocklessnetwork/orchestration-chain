

.phone: all
all:

.phony: submit-order
submit-order:
	@echo "submitting order"
	go run cmd/blockless-chaind/main.go tx market submit-order bafybeiek5cffe7cjzbmxg22maqkrms4ta6mgj5sixpcsm6hghd7kvgbd6m 10000ubls --from alice --chain-id bls-edgenet-1 --keyring-backend test --home ~/.blockless-chain -y -o json | jq