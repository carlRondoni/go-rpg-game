# base
FROM golang:1.22.0-alpine AS base

RUN adduser -S containerUser

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# build
FROM base AS build
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o /compiled ./cmd/main

# test
FROM base AS unit-test
RUN --mount=type=cache,target=/root/.cache/go-build go test -v ./cmd/main

# deploy
FROM scratch

USER containerUser

WORKDIR /

COPY --from=build /compiled /compiled
