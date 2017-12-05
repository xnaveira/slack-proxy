.PHONY: image push

IMAGE := xnaveira/slack-proxy
IMAGE_FQ := $(IMAGE):latest

image:
	docker build -t $(IMAGE_FQ) .

push: image
	docker push $(IMAGE_FQ)
