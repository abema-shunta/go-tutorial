FROM golang:1.10.2-alpine3.7

COPY src/api /go/src/api

WORKDIR /go/src/api/

ENV GIT_TERMINAL_PROMPT=1

RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

RUN apk update \
  && apk add --no-cache git \ 
  && go get -u github.com/codegangsta/gin \ 
  && go get -u github.com/golang/dep/cmd/dep \ 
  && dep ensure 
  
CMD gin -i run 
