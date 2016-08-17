VERSION_VAR := main.VERSION
REPO_VERSION := $(shell git describe --always --dirty --tags)
GOBUILD_VERSION_ARGS := -ldflags "-X $(VERSION_VAR)=$(REPO_VERSION)"
NAME := updown
TARGET := .

setup:
	go get -v github.com/kardianos/govendor
	go get -v github.com/githubnemo/CompileDaemon

clean:
	rm -f $(NAME)

vendor: FORCE
	govendor sync

$(NAME): vendor *.go
	go fmt $(TARGET)
	go get -v $(TARGET)
	go build -v -o $(NAME) -x $(GOBUILD_VERSION_ARGS)

run: vendor *.go $(NAME)
	./$(NAME)

test: $(NAME)
	go test $(TARGET)

watch:
	CompileDaemon -color=true -build "make test run"

commit-hook:
	cp dev/commit-hook.sh .git/hooks/pre-commit

version:
	@echo $(REPO_VERSION)

FORCE:
