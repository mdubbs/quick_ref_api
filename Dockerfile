FROM golang:1.13-alpine AS base
WORKDIR /app

FROM golang:1.13-alpine as build
WORKDIR /src
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" main.go

FROM base as final
WORKDIR /app
COPY --from=build /src/main .
CMD /app/main
