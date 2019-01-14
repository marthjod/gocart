PACKAGES ?= $(shell go list ./... | grep -v /vendor/ | grep -v /tests)

.PHONY: all
all: vet lint staticcheck test

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	STATUS=0; for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || STATUS=1; done; exit $$STATUS

.PHONY: staticcheck
staticcheck:
	STATUS=0; for PKG in $(PACKAGES); do CGO_ENABLED=0 staticcheck $$PKG || STATUS=1; done; exit $$STATUS

.PHONY: test
test:
	STATUS=0; for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || STATUS=1; done; exit $$STATUS

