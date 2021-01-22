FROM golang:1.15 as build
LABEL maintainer="117503445"
ENV GO111MODULE=on
WORKDIR /root/project
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app
FROM alpine as prod
EXPOSE 80
COPY --from=build /root/project/app /root/app
WORKDIR /root
ENTRYPOINT /root/app