BINARY_NAME := ovcs
SRCS := $(shell git ls-files '*.go')
LDFLAGS := "-X main.Version=$(shell git describe --tags --abbrev=0 --always) -X main.Revision=$(shell git rev-parse --verify --short HEAD)"

all: build

build: $(BINARY_NAME)

$(BINARY_NAME): $(SRCS)
	go build -ldflags $(LDFLAGS)

install:
	go install -ldflags $(LDFLAGS)

clean:
	rm -f $(BINARY_NAME)
