VERSION = latest
IMAGE = luizgustavoob/golang-encurtador-url:$(VERSION)

.PHONY: image
image:
	docker image build -t $(IMAGE) --target=image -f Dockerfile .

.PHONY: up
up: image
	docker container run --rm -it --name encurtador-url -p 8080:8080 -d $(IMAGE)