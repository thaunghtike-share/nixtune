FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y ca-certificates git

RUN rm -rf /tmp/* /var/tmp/*

ENV ACKSIN_ENV production

ADD acksin /acksin
COPY website /website

EXPOSE 8080

ENTRYPOINT /acksin server
