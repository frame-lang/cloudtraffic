FROM golang:1.17.1-alpine3.14
RUN mkdir /cloudtraffic_v2
ADD . /cloudtraffic_v2
ADD ./pubsub-system-key.json pubsub-system-key.json
ENV GOOGLE_APPLICATION_CREDENTIALS=pubsub-system-key.json
WORKDIR /cloudtraffic_v2
RUN go build -o main
CMD ["/cloudtraffic_v2/main"]