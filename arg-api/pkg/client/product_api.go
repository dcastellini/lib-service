package client

import (
	"context"
	"fmt"
	"github.com/dcastellini/lib-service/pkg/domain"
	"net/url"
)

const (
	productEndpoint     = "/api/v1/product"
	productIDQueryParam = "productID"
)

func (c *client) EditProduct(ctx context.Context, request domain.EditProductRequest) (domain.EditProductResponse, error) {
	//log := ulog.Context(ctx)
	//log.With(ulog.Any("Body", request)).Info("Logging Request Body For billing-api")

	urlPath, err := url.JoinPath(c.APIBaseUrl, productEndpoint, request.BillReminderID)
	if err != nil {
		return domain.EditProductResponse{}, fmt.Errorf("EditProduct() url %w", err)
	}

	editProductResponse := new(domain.EditProductResponse)

	resp, err := c.put(ctx, urlPath, request, editProductResponse)
	if err != nil {
		return domain.EditProductResponse{}, fmt.Errorf("request EditProduct() %w", err)
	}

	if !resp.IsSuccess() {
		return *editProductResponse, fmt.Errorf("unsuccessful EditProduct() %v", resp.Status())
	}

	return *editProductResponse, nil
}

func (c *client) GetProducts(ctx context.Context, externalClientID string) (domain.GetProductsResponse, error) {
	//log := ulog.Context(ctx)
	//log.With(ulog.Any("External-Client-ID", externalClientID)).Info("Logging Request queryStringParameter For billing-api")

	urlPath, err := url.JoinPath(c.APIBaseUrl, productEndpoint)
	if err != nil {
		return domain.GetProductsResponse{}, fmt.Errorf("GetProductsResponse() url %w", err)
	}

	getProductsResponse := new(domain.GetProductsResponse)

	resp, err := c.get(ctx, setQueryParams(productIDQueryParam, externalClientID), urlPath, getProductsResponse)
	if err != nil {
		return domain.GetProductsResponse{}, fmt.Errorf("request GetProducts() %w", err)
	}

	if !resp.IsSuccess() {
		return *getProductsResponse, fmt.Errorf("unsuccessful GetProducts() %v", resp.Status())
	}

	return *getProductsResponse, nil
}

func (c *client) CreateProduct(ctx context.Context, request domain.CreateProductRequest) (domain.CreateProductResponse, error) {
	//log := ulog.Context(ctx)
	//log.With(ulog.Any("Body", request)).Info("Logging Request Body For billing-api")

	urlEndpoint, err := url.JoinPath(c.APIBaseUrl, productEndpoint)
	if err != nil {
		return domain.CreateProductResponse{}, fmt.Errorf("CreateProduct() url %w", err)
	}

	createProductResponse := new(domain.CreateProductResponse)

	resp, err := c.post(ctx, urlEndpoint, request, createProductResponse)
	if err != nil {
		return domain.CreateProductResponse{}, fmt.Errorf("request CreateProduct() %w", err)
	}

	if !resp.IsSuccess() {
		return *createProductResponse, fmt.Errorf("unsuccessful CreateProduct() %v", resp.Status())
	}

	return *createProductResponse, nil
}

func (c *client) DeleteProduct(ctx context.Context, request domain.DeleteProductRequest) (domain.DeleteProductResponse, error) {
	//log := ulog.Context(ctx)
	//log.With(ulog.Any("Body", request)).Info("Logging Request Body For billing-api")

	urlEndpoint, err := url.JoinPath(c.APIBaseUrl, productEndpoint)
	if err != nil {
		return domain.DeleteProductResponse{}, fmt.Errorf("DeleteProduct() url %w", err)
	}

	deleteProductResponse := new(domain.DeleteProductResponse)

	resp, err := c.delete(ctx, urlEndpoint, request, deleteProductResponse)
	if err != nil {
		return domain.DeleteProductResponse{}, fmt.Errorf("request DeleteProduct() %w", err)
	}

	if !resp.IsSuccess() {
		return *deleteProductResponse, fmt.Errorf("unsuccessful DeleteProduct() %v", resp.Status())
	}

	return *deleteProductResponse, nil
}

func setQueryParams(queryParam string, value string) url.Values {
	urlValues := url.Values{}
	urlValues.Set(queryParam, value)
	return urlValues
}
