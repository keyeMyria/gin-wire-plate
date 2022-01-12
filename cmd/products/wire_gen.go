// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/products"
	"github.com/wahello/gin-wire-plate/internal/app/products/controllers"
	"github.com/wahello/gin-wire-plate/internal/app/products/grpcclients"
	"github.com/wahello/gin-wire-plate/internal/app/products/services"
	"github.com/wahello/gin-wire-plate/internal/pkg/app"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/consul"
	"github.com/wahello/gin-wire-plate/internal/pkg/jaeger"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/grpc"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/http"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	productsOptions, err := products.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := grpc.NewClientOptions(viper, tracer)
	if err != nil {
		return nil, err
	}
	client, err := grpc.NewClient(consulOptions, clientOptions)
	if err != nil {
		return nil, err
	}
	detailsClient, err := grpcclients.NewDetailsClient(client)
	if err != nil {
		return nil, err
	}
	ratingsClient, err := grpcclients.NewRatingsClient(client)
	if err != nil {
		return nil, err
	}
	reviewsClient, err := grpcclients.NewReviewsClient(client)
	if err != nil {
		return nil, err
	}
	productsService := services.NewProductService(logger, detailsClient, ratingsClient, reviewsClient)
	productsController := controllers.NewProductsController(logger, productsService)
	initControllers := controllers.CreateInitControllersFn(productsController)
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	apiClient, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, apiClient)
	if err != nil {
		return nil, err
	}
	application, err := products.NewApp(productsOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, grpcclients.ProviderSet, controllers.ProviderSet, services.ProviderSet, products.ProviderSet)
