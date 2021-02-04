.PHONY: all
all: install bot

.PHONY: install
install:
	go get -u golang.org/x/lint/golint

.PHONY: bot
bot:
	go run examples/ding-dong-bot.go

.PHONY: test
test:
	go build -o examples/ding-dong-bot -v ./examples/ding-dong-bot.go

.PHONY: clean
clean:
	rm -f examples/ding-dong-bot
