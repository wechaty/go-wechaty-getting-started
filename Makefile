.PHONY: all
all: install bot

.PHONY: install
install:
	go get -u golang.org/x/lint/golint

.PHONY: bot
bot:
	go run examples/ding-dong-bot.go