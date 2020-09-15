FROM golang AS builder
WORKDIR /root/
COPY . src
ENV GOPROXY https://goproxy.io,direct
RUN cd src && CGO_ENABLED=0 GOOS=linux go build -o ding-dong-bot -v ./examples/ding-dong-bot.go

FROM alpine AS prod
WORKDIR /root/
COPY --from=builder /root/src/ding-dong-bot .
ENTRYPOINT ["./ding-dong-bot"]
