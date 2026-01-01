.PHONY: all lint generate format

all: format lint generate

gen: generate

fmt: format

lint:
	buf lint

generate:
	buf generate

format:
	buf format -w .

