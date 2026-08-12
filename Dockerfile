# syntax=docker/dockerfile:1

FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go build

EXPOSE 80

CMD ["/envoy-test"]