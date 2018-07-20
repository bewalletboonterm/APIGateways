FROM golang:latest

ARG app_env
ENV PORT 8080
ENV APP_ENV $app_env

COPY ./app /go/src/APIGateways/app
WORKDIR /go/src/APIGateways/app

RUN go get .
RUN go build

CMD [ "app" ]
	
EXPOSE 80