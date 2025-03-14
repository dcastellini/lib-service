package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net/url"
)

const (
	acceptHeader          = "Accept"
	applicationJSONHeader = "application/json"
	contentTypeHeader     = "Content-Type"
)

func (c *client) post(ctx context.Context, url string, body, restruct any) (*resty.Response, error) {
	//log := logger.GetLogger(ctx)

	res, errPost := c.httpClient.R().
		SetContext(ctx).
		SetHeader(acceptHeader, applicationJSONHeader).
		SetHeader(contentTypeHeader, applicationJSONHeader).
		SetBody(body).
		Post(url)
	if errPost != nil {
		return nil, errPost
	}

	if errUnmarshal := c.httpClient.JSONUnmarshal(res.Body(), restruct); errUnmarshal != nil {
		//log.Debugf("billing-api ERROR unmarshaling response ", errUnmarshal)
		return res, errUnmarshal
	}

	return res, nil
}

func (c *client) get(ctx context.Context, queryParams url.Values, url string, restruct any) (*resty.Response, error) {
	//log := logger.GetLogger(ctx)

	res, errGet := c.httpClient.R().
		SetContext(ctx).
		SetHeader(acceptHeader, applicationJSONHeader).
		SetHeader(contentTypeHeader, applicationJSONHeader).
		SetQueryParamsFromValues(queryParams).
		Get(url)
	if errGet != nil {
		return nil, errGet
	}

	if errUnmarshal := c.httpClient.JSONUnmarshal(res.Body(), restruct); errUnmarshal != nil {
		//log.Debugf("billing-api ERROR unmarshaling response ", errUnmarshal)
		return res, errUnmarshal
	}

	return res, nil
}

func (c *client) put(ctx context.Context, url string, body, restruct any) (*resty.Response, error) {
	//log := logger.GetLogger(ctx)

	res, errPut := c.httpClient.R().
		SetContext(ctx).
		SetHeader(acceptHeader, applicationJSONHeader).
		SetHeader(contentTypeHeader, applicationJSONHeader).
		SetBody(body).
		Put(url)
	if errPut != nil {
		return nil, errPut
	}

	if errUnmarshal := c.httpClient.JSONUnmarshal(res.Body(), restruct); errUnmarshal != nil {
		//log.Debugf("billing-api ERROR unmarshaling response ", errUnmarshal)
		return res, errUnmarshal
	}

	return res, nil
}

func (c *client) delete(ctx context.Context, url string, body, restruct any) (*resty.Response, error) {
	//log := logger.GetLogger(ctx)

	res, errDelete := c.httpClient.R().
		SetContext(ctx).
		SetHeader(acceptHeader, applicationJSONHeader).
		SetHeader(contentTypeHeader, applicationJSONHeader).
		SetBody(body).
		Delete(url)
	if errDelete != nil {
		return nil, errDelete
	}

	if errUnmarshal := c.httpClient.JSONUnmarshal(res.Body(), restruct); errUnmarshal != nil {
		//log.Debugf("billing-api ERROR unmarshaling response ", errUnmarshal)
		return res, errUnmarshal
	}

	return res, nil
}
