FROM golang:1.17.1-alpine3.14
RUN mkdir /cloudtraffic_v1
ADD . /cloudtraffic_v1
WORKDIR /cloudtraffic_v1
RUN go build -o main
CMD ["/cloudtraffic_v1/main"]