# check system
ifeq ($(OS),Windows_NT)
	# Windows
	bin_name=qianyu-openstage.exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		# Linux
		bin_name=qianyu-openstage
	endif
	ifeq ($(UNAME_S),Darwin)
		# Mac OS X
		bin_name=qianyu-openstage
	endif
endif

build:
	@echo "build $(bin_name)..."

	go run -mod=mod github.com/99designs/gqlgen generate .
	go build -o $(bin_name) main.go

run:
	./$(bin_name)

test:
	@go test

check:
	@go vet ./
	@go fmt ./

clean:
	@go clean
