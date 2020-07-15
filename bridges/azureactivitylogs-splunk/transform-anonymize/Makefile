IMAGE := gcr.io/triggermesh/bridge-examples/azureactivitylogs-splunk-transform-anonymize

.PHONY: image release

image:
	@docker image build -t $(IMAGE) ./

release: image
	@docker image push $(IMAGE)
