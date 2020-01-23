# Selecting Base Image: i.e. golang :alpine
FROM golang:1.13.0-alpine

# Selecting the working directory
#WORKDIR $HOME/go/src/hello-world

WORKDIR $GOPATH/simple-web-server

# Copy all files from local system to alpine
COPY . .

#Installing figlet
RUN apk add figlet

#building go app
RUN go build

RUN ls -l

CMD ["./simple-web-server"]

RUN figlet GO