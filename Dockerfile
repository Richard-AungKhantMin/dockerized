FROM golang:1.23.2

LABEL maintainer="A random dude from Gritlab"
LABEL version="1.0" 
LABEL description="Docker image for Ascii-Art-Web"

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o /ascii-art-web

EXPOSE 8080

CMD ["/ascii-art-web"]
