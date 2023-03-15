TARGET = cuckoo

ifeq ($(OS),Windows_NT)
TARGET := $(TARGET).exe
endif

ifeq (n$(CGO_ENABLED),n)
CGO_ENABLED := 1
endif

RELEASE_ROOT = release

all: fmt build

build:
	@go mod download
	@echo Build paopao-ce
	@go build -trimpath -o $(RELEASE_ROOT)/$(TARGET)

run:
	@go run -trimpath -gcflags "all=-N -l" ./cmd
