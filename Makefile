SHELL=/bin/bash

build:
	npm run build
	go generate ./...
