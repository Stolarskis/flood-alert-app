FROM golang:latest

#Setup go
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/server
ADD . $GOPATH/src/github.com/server

# should be able to build now
WORKDIR $GOPATH/src/github.com/server
RUN go build && go install

EXPOSE 3001

CMD server
