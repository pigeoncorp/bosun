NAME ?= bosun-watcher
VERSION ?= $(shell cat ./VERSION)
REGISTRY ?= localhost
IMAGE ?= $(REGISTRY)/$(NAME):$(VERSION)

.PHONY: build
build:
	docker build . -t $(IMAGE)

.PHONY: publish
publish: build
	docker push $(IMAGE)