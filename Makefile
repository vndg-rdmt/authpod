target=authpod


compile:
	go build -o ./bin/$(target) ./cmd/$(target)/*.go

all: compile
