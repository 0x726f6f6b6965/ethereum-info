FROM golang:1.21.5

WORKDIR /app/project

COPY ./api/graph-service ./

RUN  rm -rf ./Dockerfile

WORKDIR /app/deps

COPY ./library ./library

COPY ./protos ./protos

RUN find protos -type f -name "*.proto" -delete
RUN find protos -type f -name "*.yaml" -delete
RUN find protos -type f -name "*.lock" -delete

WORKDIR /app/project

RUN go mod edit -replace github.com/0x726f6f6b6965/ethereum-info/library=../deps/library
RUN go mod edit -replace github.com/0x726f6f6b6965/ethereum-info/protos=../deps/protos

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./app/server.go

CMD ["/server"]