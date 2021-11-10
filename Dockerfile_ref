#1. ubuntu 설치
FROM golang:latest
MAINTAINER gun@gun.com
RUN apt-get -y update

#2. go 소스 실행
RUN apt-get install -y git
RUN git clone https://github.com/grapegun67/go1.git
RUN go run go1/test.go

#3. protoc 설치(에러가 난다면 protoc 버전을 최신으로 해줄 것)
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
RUN apt-get install -y unzip
RUN unzip protoc-3.19.1-linux-x86_64.zip -d $HOME/.local
RUN export PATH=$PATH:$HOME/.local/bin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN export PATH=$PATH:$(go env GOPATH)/bin
