FROM golang:1.22 as builder
RUN mkdir /build
COPY . /build
WORKDIR /build

ARG SSH_PRIVATE_KEY
ARG VERSION
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/*

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.version=${VERSION}'" -a -o app ./cmd

FROM alpine:3.19
RUN apk update && apk upgrade -U -a
COPY --from=builder /build/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
