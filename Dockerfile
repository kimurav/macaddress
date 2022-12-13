FROM golang:alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /brightsign
EXPOSE 8080
ENTRYPOINT [ "/brightsign" ]