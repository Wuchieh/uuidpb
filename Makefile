.PHONY: all lint generate format

all: format lint generate

gen: generate

fmt: format

lint:
	buf lint

generate:
	buf generate
	mv gen/uuidpb/v1/uuid.pb.go .

format:
	buf format -w .

