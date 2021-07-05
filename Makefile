VERSION = latest
IMAGE = luizgustavoob/golang-encurtador-url:$(VERSION)

.PHONY: build
build:
	docker image build -t $(IMAGE) --target=build -f Dockerfile .

.PHONY: image
image:
	docker image build -t $(IMAGE) --target=image -f Dockerfile .

.PHONY: up
up: image
	docker container run --rm -it --name encurtador-url -p 8888:8888 -d $(IMAGE)