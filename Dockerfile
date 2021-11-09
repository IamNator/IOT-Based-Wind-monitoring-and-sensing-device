# Start from golang base image
FROM golang:1.17-alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="iamnator <natorverinumbe@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk --no-cache add tzdata
RUN go version

# creates working directory for program
WORKDIR /go/src/github.com/IamNator/iot

# copies all program files specified directory in the container
ADD . .



#RUN go mod download
RUN go get .
RUN go get github.com/pilu/fresh
RUN chmod +x scripts/startapp.sh

CMD ["./scripts/startapp.sh"]

EXPOSE ${PORT}