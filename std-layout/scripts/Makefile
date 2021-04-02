git_rev    = $(shell git rev-parse --short HEAD)
git_tag    = $(shell git describe --tags --abbrev=0)
git_branch = $(shell git rev-parse --abbrev-ref HEAD)

VERSION := $(git_branch)_$(git_rev)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# in detached HEAD state
ifeq ($(git_branch), HEAD)
	git_branch = $(shell git show-ref | grep $(shell git show HEAD | sed -n 1p | cut -d " " -f 2) | sed 's|.*/\(.*\)|\1|' | grep -v HEAD | sort | uniq | head -n 1)
	# when git checkout <<tag>>, branch may still be empty
	ifeq ($(git_branch), )
		git_branch := $(git_tag)
	endif
	VERSION := $(git_branch)_$(git_rev)
endif

ifeq ($(git_branch), develop)
	VERSION := develop_$(git_rev)
endif

ifeq ($(git_branch), master)
	VERSION := release_$(git_tag)_$(git_rev)
endif

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.BuildTime=$(BUILDTIME)

build:
	@go build -o atlas-acs -ldflags "$(GOLDFLAGS)" .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o atlas-acs -ldflags "$(GOLDFLAGS)" .
