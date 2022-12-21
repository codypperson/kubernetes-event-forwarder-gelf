FROM golang:1.16
COPY . /src/
WORKDIR /src/
RUN make clean \
  && make test \
  && make

FROM ubuntu:focal
COPY --from=0 /src/event-forwarder-gelf /event-forwarder-gelf
ENTRYPOINT ["/event-forwarder-gelf"]
