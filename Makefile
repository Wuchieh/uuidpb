.PHONY: all generate format

all: format generate

gen: generate

fmt: format

generate:
	buf generate

format:
	buf format -w .

