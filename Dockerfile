FROM golang:1.23

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ditto-file-processor

CMD [ "ditto-file-processor ]