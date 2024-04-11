FROM golang:latest

ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait_postgres.sh

# build service
RUN go mod download
RUN go build -o banner-server ./cmd/banner-server/main.go 

CMD ["./banner-server"]
