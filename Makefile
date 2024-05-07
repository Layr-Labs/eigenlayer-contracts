
.PHONY: install-hooks
install-hooks:
	cp scripts/pre-commit.sh .git/hooks/pre-commit

.PHONY: deps
deps: install-hooks
	brew install libusb
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	curl -L https://foundry.paradigm.xyz | bash
	foundryup

.PHONY: compile
compile:
	forge b

.PHONY: bindings
bindings: compile
	./scripts/compile-bindings.sh
