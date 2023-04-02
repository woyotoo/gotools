DATE=`date +%FT%T%z`
GO_VERSION=`go version`

ifeq (, $(shell which go))
$(warning "=> check: could not find go command.")
endif

all: info
	@echo '[I] Please run test:'
	@echo '    make test'
test: info
	@echo "[I] make test start ..."
	@go test -cpu=1,2,4 -cover --coverprofile coverage.out ./...
	@echo "[I] make test done."

test2:
	@echo "[I] make test start ..."
	@go test -cpu=1,2,4 -v -tags integration -cover --coverprofile coverage.out ./...
	@echo "[I] make test done."

info:
	@echo "[I] Date: $(DATE)"
	@echo "[I] check: $(GO_VERSION)"