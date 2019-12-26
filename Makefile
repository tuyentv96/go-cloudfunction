.DEFAULT_GOAL := all

.PHONY: default test

deploy:
	gcloud functions deploy Handler --runtime=go111  --trigger-http