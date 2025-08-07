package envconfig

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type GrpcClientConfig struct {
	GrpcServerHost string `env:"GRPC_SERVER_HOST, required"`
	GrpcServerPort string `env:"GRPC_SERVER_PORT, default=443"`
}

func (cfg GrpcClientConfig) GrpcServerAddress() string {
	return cfg.GrpcServerHost + ":" + cfg.GrpcServerPort
}

type RunConfig struct {
	JobId           int    `env:"JOB_ID, required"`
	ChangeRequestId int    `env:"CHANGE_REQUEST_ID, required"`
	RunType         string `env:"RUN_TYPE, required"` // TMT or Rebase
}

func LoadGrpcClientConfig() (GrpcClientConfig, error) {
	var cfg GrpcClientConfig
	err := envconfig.Process(context.Background(), &cfg)
	return cfg, err
}

func LoadRunConfig() (RunConfig, error) {
	var cfg RunConfig
	err := envconfig.Process(context.Background(), &cfg)
	return cfg, err
}
