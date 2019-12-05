FROM golang

LABEL maintainer="Glauber <sistemas.glauber@gmail.com>"

WORKDIR /app/src/go-api-jwt-token
ENV GOPATH=/app
COPY . /app/src/go-api-jwt-token
RUN go build main.go
ENTRYPOINT ./main
EXPOSE 8080

