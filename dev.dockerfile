FROM golang:alpine AS builder
WORKDIR /app
COPY ./src .
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN apk add --update git && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g handlers/routes.go
RUN air init    
CMD [ "air", "--build.args_bin", "serve"]
