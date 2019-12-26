.DEFAULT_GOAL := all

.PHONY: default test

codegen:
	gcloud functions deploy Handler --runtime=go111  --trigger-http