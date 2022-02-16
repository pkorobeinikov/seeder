# Improve with build stage
FROM golang:1.17.6

WORKDIR /usr/src/seeder

COPY go.mod go.sum ./

RUN go mod download -x && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/seeder ./...

ENTRYPOINT ["seeder"]
