GOFUMPT_BIN=$(LOCAL_BIN)/gofumpt
$(GOFUMPT_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install mvdan.cc/gofumpt@v0.6.0

IMPREVISER_BIN=$(LOCAL_BIN)/goimports-reviser
$(IMPREVISER_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/incu6us/goimports-reviser/v3@v3.6.4

GOLINT_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLINT_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61

GOOSE_BIN=$(LOCAL_BIN)/goose
$(GOOSE_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/pressly/goose/cmd/goose@v2.7.0