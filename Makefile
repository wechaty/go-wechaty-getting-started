.PHONY: all
all: install bot

.PHONY: install
install:
	go install golang.org/x/lint/golint@latest

.PHONY: bot
bot:
	go run examples/ding-dong-bot.go

.PHONY: test
test:
	go build -o examples/ding-dong-bot -v ./examples/ding-dong-bot.go
	go build -o examples/plugln/ding-dong-bot -v ./examples/plugln/ding-dong-bot.go

.PHONY: clean
clean:
	rm -f examples/ding-dong-bot
	rm -f examples/plugln/ding-dong-bot
