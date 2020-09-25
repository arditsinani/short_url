FROM golang:latest as dev
LABEL maintainer="Ardit Sinani <arditsinani04@gmail.com>"

# copy the local package files to the container workspace
ADD . /go/src/short_url/app

# Setting up working directory
WORKDIR /go/src/short_url/app

# Install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Listen to directory changes
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build ./main.go" --command=./main