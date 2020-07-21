FROM golang
COPY . go_wechaty_examples
ENV GOPROXY https://goproxy.io,direct
RUN cd go_wechaty_examples && make test
ENTRYPOINT ./go_wechaty_examples/ding-dong-bot