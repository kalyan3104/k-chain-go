FROM golang:1.20.7 as builder

RUN apt-get update && apt-get install -y
WORKDIR /go/k-chain-go
COPY . .
RUN go mod tidy
# Seed node
WORKDIR /go/k-chain-go/cmd/seednode
RUN go build

# ===== SECOND STAGE ======
FROM ubuntu:22.04
COPY --from=builder /go/k-chain-go/cmd/seednode /go/k-chain-go/cmd/seednode

WORKDIR /go/k-chain-go/cmd/seednode/
EXPOSE 10000
ENTRYPOINT ["./seednode"]
