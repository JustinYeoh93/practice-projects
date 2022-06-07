FROM go:1.17.9-buster

WORKDIR /usr/src/app

COPY go.sum go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/...

CMD ["app"]