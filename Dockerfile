FROM golang:alpine

WORKDIR /usr/local/go/src/github.com/virtual-kubelet/virtual-kubelet
COPY . .
ENV PATH=$PATH:/usr/local/go/bin
ENV HOME=/root
ENV GOPATH=/usr/local/go
ENV GO111MODULE=off

RUN apk update
RUN apk add git && apk add ca-certificates
RUN CGO_ENABLED=0 go build -tags wlm_provider -o main cmd/virtual-kubelet/main.go

ENTRYPOINT [ "./main" ]
