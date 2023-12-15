name = "block-service"
grpc {
    host = "localhost"
    port = 64535
    network = "tcp"
}

db {
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "docker"
    db-name = "postgres"
}

redis {
    host = "localhost"
    port = 6379
    password = "pwd"
    db = 0
    max-retries = 3
}

rpc {
    servers = ["https://data-seed-prebsc-2-s3.binance.org:8545"]
    connect-timeout = 30
    retry = 3
}

log {
    service-name = "block-service"
}