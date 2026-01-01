.PHONY: all lint generate format

all: format lint generate

gen: generate

fmt: format

lint:
	buf lint

generate:
	buf generate
	mv wuchieh/protobuf/* .
	rm -rf wuchieh

format:
	buf format -w .

