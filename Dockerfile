FROM golang:latest

ENV PORT 80
ENV HOST_BEWALLET localhost:9300

COPY ./app /go/src/APIGateways/app
WORKDIR /go/src/APIGateways/app

RUN go get .
RUN go build

CMD [ "app" ]
	
EXPOSE 80