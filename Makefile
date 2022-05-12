
build: GOOS=$(shell go env GOOS)
build: GOARCH=$(shell go env GOARCH)
build: DESTINATION=$(HOME)/.terraform.d/plugins/mlin.com/mlin/mlin/99.0.0/$(GOOS)_$(GOARCH)
build: ${mkdir -p $(DESTINATION)}
	go build -o $(DESTINATION)/terraform-provider-mlin_v99.0.0