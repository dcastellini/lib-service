package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type APIConfiguration struct {
	API struct {
		UserAgentHeader     string `env:"USER_AGENT" env-default:"uala lambda" env-description:"User Agent"`
		TransportName       string `env:"TRANSPORT_NAME" env-default:"billing-api" env-description:"Transport Name"`
		TCPTimeOut          int64  `env:"TCP_TIMEOUT" env-default:"80" env-description:"TCP Time Out"`
		KeepAliveProbesTime int64  `env:"KEEPALIVE_PROBES_TIMEOUT" env-default:"80" env-description:"Keep Alive Probes Time"`
		MaxIdleConns        int    `env:"MAX_IDLE_CONNS" env-default:"100" env-description:"Max Idle Conns"`
		MaxConnsPerHost     int    `env:"MAX_CONNS_PER_HOST" env-default:"100" env-description:"Max Conns Per Host"`
		MaxIdleConnsPerHost int    `env:"MAX_IDLE_CONNS_PER_HOST" env-default:"100" env-description:"Max Idle Conns Per Host"`
		Timeout             int64  `env:"TIMEOUT_SEC" env-default:"35" env-description:"Timeout in seconds"`
		ProductAPIBaseURL   string `env:"BASE_URL" env-default:"https://api.test.amazonaws.com" env-description:"Base URL"`
		RetryWaitTime       int64  `env:"RETRY_WAIT_TIME" env-default:"1" env-description:"Retry Wait Time in seconds"`
		RetryMaxWaitTime    int64  `env:"RETRY_MAX_WAIT_TIME" env-default:"2" env-description:"Retry Max Wait Time in seconds"`
		RetryCount          uint   `env:"RETRY_COUNT" env-default:"0" env-description:"Retry Count"`
	}
	General struct {
		Environment string `env:"ENVIRONMENT" env-default:"local" env-description:"Lambda Environment"`
		Region      string `env:"AWS_REGION" env-default:"us-east-1" env-description:"AWS Region"`
		Country     string `env:"COUNTRY" env-default:"ARG" env-description:"ARG MEX COL MEXBANK (MULTI)"`
	}
	RoundTripperMiddleware struct {
		EnabledDumpHTTP bool `env:"DUMP_HTTP" env-default:"true" env-description:"Enabled HTTP DUMP"`
		Enabled         bool `env:"HTTP_LOG" env-default:"true" env-description:"Enabled HTTP LOGGING"`
	}
}

// NewAPIConfiguration returns filled LambdaConfiguration.
func NewAPIConfiguration() *APIConfiguration {
	cfg := &APIConfiguration{}

	return cfg
}

func (cfg *APIConfiguration) LoadFromEnvs() {
	if err := cleanenv.ReadEnv(cfg); err != nil {
		panic(err)
	}
}

// GetEnvsDescriptions get envs description.
func (cfg *APIConfiguration) GetEnvsDescriptions() string {
	header := "Environment variables"
	help, _ := cleanenv.GetDescription(cfg, &header)

	return help
}
