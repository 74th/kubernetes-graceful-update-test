FROM golang:1.13 AS build-env

WORKDIR /work
COPY main.go ./
RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /work/server .
EXPOSE 8080
ENTRYPOINT ./server
