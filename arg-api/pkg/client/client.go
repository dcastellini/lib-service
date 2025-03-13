package client

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

func NewAPIClient(opts ...Options) API {
	cfg := config.NewAPIConfiguration()
	cfg.LoadFromEnvs()

	apiClient := &client{
		httpClient:              newRestyClient(cfg),
		transportName:           cfg.API.TransportName,
		enabledLoggerMiddleware: cfg.RoundTripperMiddleware.Enabled,
		dumpHTTP:                cfg.RoundTripperMiddleware.EnabledDumpHTTP,
		billingAPIBaseURL:       cfg.BillingAPI.BillingAPIBaseURL,
		cfg:                     cfg,
	}

	for _, opt := range opts {
		opt(apiClient)
	}

	return apiClient
}

func newRestyClient(cfg *config.APIConfiguration) *resty.Client {
	httpClient := defaultHTTPClient(cfg)

	restyClient := resty.NewWithClient(httpClient)
	restyClient.JSONMarshal = json.Marshal
	restyClient.JSONUnmarshal = json.Unmarshal

	restyClient.SetTransport(defaultHTTPLogger(restyClient.GetClient().Transport, cfg))

	restyClient.SetRetryCount(int(cfg.BillingAPI.RetryCount)).
		SetRetryWaitTime(time.Duration(cfg.BillingAPI.RetryWaitTime) * time.Second).
		SetRetryMaxWaitTime(time.Duration(cfg.BillingAPI.RetryMaxWaitTime) * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() >= http.StatusInternalServerError
			},
		)

	return restyClient
}

func defaultHTTPClient(cfg *config.APIConfiguration) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()

	transport.DialContext = func(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
		return dialer.DialContext
	}(&net.Dialer{
		Timeout:   time.Duration(cfg.BillingAPI.TCPTimeOut) * time.Second,
		KeepAlive: time.Duration(cfg.BillingAPI.KeepAliveProbesTime) * time.Second,
	})

	transport.MaxIdleConns = cfg.BillingAPI.MaxIdleConns
	transport.MaxConnsPerHost = cfg.BillingAPI.MaxConnsPerHost
	transport.MaxIdleConnsPerHost = cfg.BillingAPI.MaxIdleConnsPerHost

	return &http.Client{
		Timeout:   time.Duration(cfg.BillingAPI.Timeout) * time.Second,
		Transport: transport,
	}
}

func defaultHTTPLogger(transport http.RoundTripper, cfg *config.APIConfiguration) http.RoundTripper {
	if cfg.RoundTripperMiddleware.Enabled {
		return middleware.NewLoggingRoundTripper(
			transport,
			cfg.BillingAPI.TransportName,
			cfg.RoundTripperMiddleware.EnabledDumpHTTP,
		)
	}

	return transport
}
