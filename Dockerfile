FROM golang:alpine

COPY gogogadget /opt/gogogadget/

COPY views /opt/gogogadget/views

WORKDIR /opt/gogogadget

CMD ["/opt/gogogadget/gogogadget"]