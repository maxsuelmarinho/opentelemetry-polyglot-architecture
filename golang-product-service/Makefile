#!make

build_directory=./build
app_name=server
home				        = 	$(shell home)
app_version =	$(shell cat VERSION)
version_array = $(subst ., ,$(app_version))
major =	$(word 1,${version_array})
minor = $(word 2,${version_array})
patch = $(word 1,$(subst -, ,$(word 3,${version_array})))
pre_release_build_array =	$(subst ., , $(subst -, ,$(word 3,${version_array})))
pre_release_build =	$(or $(word 4,${version_array}), -1)
pwd 				        = 	$(shell pwd)

help:				## Show this help.
	- @fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# Build targets
patch:
	- @echo "BUMPING PATCH"
	- @echo "Current Version: $(app_version)"
	- $(eval patch=$(shell echo $$(($(patch)+1))))
	- @echo "New Version: $(major).$(minor).$(patch)"
	- @printf $(major).$(minor).$(patch) > VERSION

minor:
	- @echo "BUMPING MINOR"
	- @echo "Current Version: $(app_version)"
	- $(eval minor=$(shell echo $$(($(minor)+1))))
	- @echo "New Version: $(major).$(minor).0"
	- @printf $(major).$(minor).0 > VERSION

major:
	- @echo "BUMPING MAJOR"
	- @echo "Current Version: $(app_version)"
	- $(eval major=$(shell echo $$(($(major)+1))))
	- @echo "New Version: $(major).0.0"
	- @printf $(major).0.0 > VERSION

pre-release:
	- @echo "BUMPING PRE-RELEASE BUILD"
	- @echo "Current Version: $(app_version)"
	- $(eval pre_release_build=$(shell echo $$(($(pre_release_build)+1))))
	- @echo "New Version: $(major).$(minor).$(patch)-rc.$(pre_release_build)"
	- @printf $(major).$(minor).$(patch)-rc.$(pre_release_build) > VERSION

promote:
	- @echo "PROMOTING PRE-RELEASE"
	- @echo "Current Version: $(app_version)"
	- @echo "New Version: $(major).$(minor).$(patch)"
	- @printf $(major).$(minor).$(patch) > VERSION

display-version:		## Display current version
	- @echo "v$(app_version)"

build: 			## Generate application binaries
	@go mod download && \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -installsuffix cgo -o ${build_directory}/${app_name} cmd/server/main.go && \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -installsuffix cgo -o ${build_directory}/healthchecker cmd/healthchecker/main.go || exit $?

# Quality target
install-go-tools:
# workaround to globally install binaries without adding it to go.mod. see open issue thread: https://github.com/golang/go/issues/30515
	- cd ~/
	- go get github.com/hhatto/gocloc/cmd/gocloc
	- go get golang.org/x/lint/golint
	- cd -
	- go mod tidy

test: 			## Execute application unit tests
	@go test -v ./...

ifeq (test cover,$(MAKECMDGOALS))
cover:
	@go test -coverprofile=coverage.out -v ./...
	@go tool cover -html=coverage.out -o=coverage.html
endif

integration-test: 	## Execute application integration tests
	@go test -v ./... -tags=integration -run Integration

lint: install-go-tools		## Execute static code analysis
	@go vet ./...
	@golint -set_exit_status ./... || exit $?
	@if [ "$(shell gofmt -s -l ./ | wc -l)" -gt 0 ]; then echo "code is not formatted properly" && exit 1; fi
	@gocloc .
	@go mod tidy

validate: build test lint

.PHONY: build
