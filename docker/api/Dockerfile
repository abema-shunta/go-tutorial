FROM golang:1.10.2-alpine3.7

COPY src/api /go/src/api

WORKDIR /go/src/api/

ENV GIT_TERMINAL_PROMPT=1

RUN apk update \
  && apk add --no-cache git

RUN git config --global url.'git@github.com'.insteadof https://github.com

RUN go get -u github.com/codegangsta/gin.git \ 
  && go get -u github.com/golang/dep/cmd/dep.git \ 
  && dep ensure 
  
CMD gin -i run 
