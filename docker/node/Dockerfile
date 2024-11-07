FROM golang:1.20.7 as builder

RUN apt-get update && apt-get upgrade -y
WORKDIR /go/k-chain-go
COPY . .
RUN go mod tidy
# Kalyan node
WORKDIR /go/k-chain-go/cmd/node
RUN go build -v -ldflags="-X main.appVersion=$(git describe --tags --long --dirty)"

RUN cp /go/pkg/mod/github.com/kalyan3104/$(cat /go/k-chain-go/go.mod | grep k-chain-vm-v | sort -n | tail -n -1 | awk -F '/' '{print$3}' | sed 's/ /@/g')/wasmer/libwasmer_linux_amd64.so /lib_amd64/
RUN cp /go/pkg/mod/github.com/kalyan3104/$(cat /go/k-chain-go/go.mod | grep k-chain-vm-go | sort -n | tail -n -1 | awk -F '/' '{print$3}' | sed 's/ /@/g')/wasmer2/libvmexeccapi.so /lib_amd64/

RUN cp /go/pkg/mod/github.com/kalyan3104/$(cat /go/k-chain-go/go.mod | grep k-chain-vm-v | sort -n | tail -n -1 | awk -F '/' '{print$3}' | sed 's/ /@/g')/wasmer/libwasmer_linux_arm64_shim.so /lib_arm64/
RUN cp /go/pkg/mod/github.com/kalyan3104/$(cat /go/k-chain-go/go.mod | grep k-chain-vm-go | sort -n | tail -n -1 | awk -F '/' '{print$3}' | sed 's/ /@/g')/wasmer2/libvmexeccapi_arm.so /lib_arm64/

# ===== SECOND STAGE ======
FROM ubuntu:22.04
RUN apt-get update && apt-get upgrade -y
COPY --from=builder "/go/k-chain-go/cmd/node/node" "/go/k-chain-go/cmd/node/"

# Copy architecture-specific files
COPY --from=builder "/lib_${TARGETARCH}/*" "/lib/"

WORKDIR /go/k-chain-go/cmd/node/
EXPOSE 8080
ENTRYPOINT ["/go/k-chain-go/cmd/node/node"]
