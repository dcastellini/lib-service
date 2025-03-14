package client

import (
	"context"
	"github.com/dcastellini/lib-service/pkg/config"
	"github.com/dcastellini/lib-service/pkg/domain"
	"github.com/go-resty/resty/v2"
)

type API interface {
	IProductAPI
}

type IProductAPI interface {
	DeleteProduct(ctx context.Context, request domain.DeleteProductRequest) (domain.DeleteProductResponse, error)
	EditProduct(ctx context.Context, request domain.EditProductRequest) (domain.EditProductResponse, error)
	GetProducts(ctx context.Context, externalClientID string) (domain.GetProductsResponse, error)
	CreateProduct(ctx context.Context, request domain.EditProductRequest) (domain.EditProductResponse, error)
}

type client struct {
	httpClient              *resty.Client
	cfg                     *config.APIConfiguration
	APIBaseUrl              string
	transportName           string
	enabledLoggerMiddleware bool
	dumpHTTP                bool
}
