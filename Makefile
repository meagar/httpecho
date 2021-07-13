.PHONY: server
server:
	air -c .air.toml

.PHONY: install
install:
	go install -tags release