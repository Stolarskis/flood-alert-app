FROM golang:latest

#Setup go
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
ENV KUBERNETES_IP=10.7.240.2

RUN mkdir -p $GOPATH/src/github.com/app-flood-server
ADD . $GOPATH/src/github.com/app-flood-server

# should be able to build now
WORKDIR $GOPATH/src/github.com/app-flood-server
RUN go build && go install
