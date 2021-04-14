FROM golang

ADD . /go/src/github.com/bgoldovsky/service-dddealer

WORKDIR /go/src/github.com/bgoldovsky/service-dddealer
RUN go install /go/src/github.com/bgoldovsky/service-dddealer/cmd/service

ENTRYPOINT /go/bin/service