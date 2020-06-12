NAME:=$(shell basename `git rev-parse --show-toplevel`)
HASH:=$(shell git rev-parse --verify HEAD)

all: dockerbuild

clean:
	rm -rf pkg bin

up: dockerbuild down
	docker-compose up

dockerbuild:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(HASH)" -o $(NAME)
	docker build -t $(NAME) .

build:
	go build -ldflags "-X main.version=$(HASH)" -o $(NAME)

down:
	docker-compose down
