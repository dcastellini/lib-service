package client

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Options func(*client)

func WithCustomHTTPClient(httpClient *http.Client) Options {
	return func(c *client) {
		c.httpClient = resty.NewWithClient(httpClient)
	}
}

func WithCustomAPIBaseURL(apiBaseURL string) Options {
	return func(c *client) {
		c.APIBaseUrl = apiBaseURL
	}
}

func WithCustomTransportName(transport string) Options {
	return func(c *client) {
		c.transportName = transport
	}
}

func WithCustomHeaders(headers map[string]string) Options {
	return func(c *client) {
		c.httpClient.SetHeaders(headers)
	}
}

func WithCustomHeader(name, value string) Options {
	return func(c *client) {
		c.httpClient.SetHeader(name, value)
	}
}

func WithCustomUserAgent(userAgent string) Options {
	return func(c *client) {
		c.httpClient.SetHeader("User-Agent", userAgent)
	}
}
