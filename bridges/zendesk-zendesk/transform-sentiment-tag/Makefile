IMAGE := gcr.io/triggermesh/bridge-examples/zendesk-zendesk-transform-sentiment-tag

.PHONY: image release

image:
	@docker image build -t $(IMAGE) ./

release: image
	@docker image push $(IMAGE)
