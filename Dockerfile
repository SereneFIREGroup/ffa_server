FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/ffa_server ./cmd/...

# Path: Dockerfile
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/bin/ /app/bin/

EXPOSE 10001

ENV ENVIRONMENT=development

ENTRYPOINT ["/app/bin/"]

CMD ["./ffa_server"]