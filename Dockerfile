FROM golang:1.15.12-alpine3.13
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o nearestprime cmd/appserver/main.go
CMD [ "/app/nearestprime" ]