FROM golang:1.17.2-alpine3.14 as build
LABEL maintainer="117503445"
ENV GO111MODULE=on
WORKDIR /root/project
ADD go.mod .
ADD go.sum .
RUN go mod download
RUN wget https://github.com/gogf/gf-cli/releases/download/v1.17.0/gf_linux_amd64 && chmod +x gf && ./gf install
ADD . .
RUN gf swagger --pack -y
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o server_bin
FROM alpine:3.14 as prod
EXPOSE 8080
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
WORKDIR /root
COPY --from=build /root/project/server_bin /root/server_bin
ENTRYPOINT /root/server_bin