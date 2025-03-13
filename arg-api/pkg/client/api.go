package client

import (
	"context"
)

type API interface {
	IProductAPI
}

type IProductAPI interface {
	DeleteProduct(ctx context.Context, request domain.) (domain.EditBillReminderResponse, error)
	EditProduct(ctx context.Context, request domain.EditBillReminderRequest) (domain.EditBillReminderResponse, error)
	GetProducts(ctx context.Context, externalClientID string) (domain.GetBillRemindersResponse, error)
	CreateProduct(ctx context.Context, request domain.CreateBillReminderRequest) (domain.CreateBillReminderResponse, error)
}

type client struct {
	httpClient              *resty.Client
	cfg                     *config.APIConfiguration
	APIBaseUrl              string
	transportName           string
	enabledLoggerMiddleware bool
	dumpHTTP                bool
}
