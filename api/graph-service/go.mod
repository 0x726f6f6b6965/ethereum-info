module github.com/0x726f6f6b6965/ethereum-info/graph-service

go 1.21.5

require (
	github.com/0x726f6f6b6965/ethereum-info/library v0.0.0-00010101000000-000000000000
	github.com/0x726f6f6b6965/ethereum-info/protos v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.17.41
	github.com/joho/godotenv v1.5.1
	github.com/vektah/gqlparser/v2 v2.5.10
	go.uber.org/zap v1.26.0
	google.golang.org/grpc v1.60.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.3 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.1.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/urfave/cli/v2 v2.25.7 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/0x726f6f6b6965/ethereum-info/protos => ../../protos

replace github.com/0x726f6f6b6965/ethereum-info/library => ../../library
