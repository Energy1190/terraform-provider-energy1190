all: build

build:
	CGO_ENABLED=0 go build -o ./dist/terraform-provider-superset

.PHONY: build