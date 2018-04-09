FROM golang:1.9-alpine

RUN \
     apk update \
  && apk add curl git \
  && rm -rf /var/cache/apk/*

RUN curl https://glide.sh/get | sh

RUN mkdir -p /root/.glide

ENV GLIDE_HOME /root/.glide
