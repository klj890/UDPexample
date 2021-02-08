FROM busybox:latest

WORKDIR /UDPexample

ADD UDPexample ./

COPY ./examples ./examples/