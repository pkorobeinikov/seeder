FROM golang:1.17.8 AS builder

WORKDIR /usr/src/seeder

COPY go.mod go.sum ./

RUN go mod download -x && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -v -ldflags="-w -s" -o /bin/seeder cmd/seeder/main.go



FROM scratch

COPY --from=builder /bin/seeder /bin/seeder

ENTRYPOINT ["/bin/seeder"]
