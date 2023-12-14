package config

import "time"

type Grpc struct {
	Host       string `hcl:"host" help:"the host to bind for GRPC server"`
	Port       int    `hcl:"port" help:"the port to bind for GRPC server"`
	Network    string `hcl:"network" help:"the network type for GRPC server" default:"tcp"`
	Socketpath string `hcl:"socketpath" help:"the socket path when network is unix"`
}

type Rest struct {
	Host string `hcl:"host" help:"the host to bind for REST server"`
	Port int    `hcl:"port" help:"the port to bind for REST server"`
}

type Client struct {
	Host              string        `hcl:"host" help:"the host to connect server"`
	GrpcPort          int           `hcl:"grpc-port" help:"the port to bind for GRPC server"`
	GrpcNetwork       string        `hcl:"grpc-network" help:"the network type for GRPC server" default:"tcp"`
	GrpcSocketpath    string        `hcl:"grpc-socketpath" help:"the socket path when network is unix"`
	HTTPPort          int           `hcl:"port" help:"the port to bind for REST server"`
	ConnectionTimeout time.Duration `hcl:"connection-timeout" help:"the service connection timeout"`
}

type Log struct {
	// Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	Level            int    `hcl:"level" help:"the application log level" default:"1"`
	TimeFormat       string `hcl:"time-format" help:"the application log time format" default:"2006-01-02T15:04:05Z07:00"`
	TimestampEnabled bool   `hcl:"timestamp-enabled" default:"false"`
	ServiceName      string `hcl:"service-name" help:"the application service name"`
}

type Config struct {
	Name    string `hcl:"name" help:"the application name"`
	Rest    Rest   `hcl:"rest" help:"the application rest option"`
	Log     Log    `hcl:"log" help:"the application log"`
	Clients struct {
		Block       Client `hcl:"block-svc"`
		Transaction Client `hcl:"transaction-svc"`
	} `hcl:"clients" help:"the clients the application talk to"`
}
