FROM golang:latest

#Setup go
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   

RUN mkdir -p $GOPATH/src/github.com/app
ADD . $GOPATH/src/github.com/app

# should be able to build now
WORKDIR $GOPATH/src/github.com/app
RUN go install 

EXPOSE 3000
CMD [ "app" ]
