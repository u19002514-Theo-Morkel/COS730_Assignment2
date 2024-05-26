FROM golang:1.22.2-alpine3.18 AS Build

WORKDIR /app

COPY ./go.mod ./

RUN go mod tidy

RUN go mod download

COPY ./cmd/PageService/main.PageService.go ./cmd/PageService/main.PageService.go

COPY . ./

RUN go build -o main ./cmd/PageService/main.PageService.go

FROM alpine:3.18

COPY --from=Build /app/main /app/main

CMD ["/app/main"]

