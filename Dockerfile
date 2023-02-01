FROM golang:alpine AS builder
WORKDIR /app
COPY ./src .
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN apk add --update git
RUN go build -a -installsuffix cgo -o godemo .

FROM alpine:edge
WORKDIR /usr/src/app
COPY --from=builder /app/godemo .

RUN mv /usr/src/app/godemo /usr/bin

CMD ["godemo serve"]
