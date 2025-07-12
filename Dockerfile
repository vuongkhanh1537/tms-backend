# BUILD STAGE
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server ./cmd/server

# RUN STAGE
FROM scratch

WORKDIR /app

COPY ./config /app/config

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
