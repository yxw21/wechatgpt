FROM golang:1.19
WORKDIR /wechatgpt
COPY . .
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install gcc g++ ca-certificates chromium xvfb && \
    go mod download && \
    go build wechatgpt.go
ENTRYPOINT [ "/wechatgpt/wechatgpt" ]