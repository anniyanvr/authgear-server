// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package images

import (
	"github.com/authgear/authgear-server/pkg/images/deps"
	"github.com/authgear/authgear-server/pkg/images/handler"
	"github.com/authgear/authgear-server/pkg/images/service"
	deps2 "github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/images"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/middleware"
	"github.com/authgear/authgear-server/pkg/lib/presign"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/httputil"
	"net/http"
)

// Injectors from wire.go:

func newPanicMiddleware(p *deps.RootProvider) httproute.Middleware {
	factory := p.LoggerFactory
	panicMiddlewareLogger := middleware.NewPanicMiddlewareLogger(factory)
	panicMiddleware := &middleware.PanicMiddleware{
		Logger: panicMiddlewareLogger,
	}
	return panicMiddleware
}

func newSentryMiddleware(p *deps.RootProvider) httproute.Middleware {
	hub := p.SentryHub
	environmentConfig := &p.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	sentryMiddleware := &middleware.SentryMiddleware{
		SentryHub:  hub,
		TrustProxy: trustProxy,
	}
	return sentryMiddleware
}

func newCORSMiddleware(p *deps.RequestProvider) httproute.Middleware {
	appProvider := p.AppProvider
	config := appProvider.Config
	appConfig := config.AppConfig
	httpConfig := appConfig.HTTP
	oAuthConfig := appConfig.OAuth
	samlConfig := appConfig.SAML
	rootProvider := appProvider.RootProvider
	environmentConfig := &rootProvider.EnvironmentConfig
	corsAllowedOrigins := environmentConfig.CORSAllowedOrigins
	corsMatcher := &middleware.CORSMatcher{
		Config:             httpConfig,
		OAuthConfig:        oAuthConfig,
		SAMLConfig:         samlConfig,
		CORSAllowedOrigins: corsAllowedOrigins,
	}
	factory := rootProvider.LoggerFactory
	corsMiddlewareLogger := middleware.NewCORSMiddlewareLogger(factory)
	corsMiddleware := &middleware.CORSMiddleware{
		Matcher: corsMatcher,
		Logger:  corsMiddlewareLogger,
	}
	return corsMiddleware
}

func newGetHandler(p *deps.RequestProvider) http.Handler {
	appProvider := p.AppProvider
	rootProvider := appProvider.RootProvider
	objectStoreConfig := rootProvider.ObjectStoreConfig
	clock := _wireSystemClockValue
	imagesCloudStorageServiceStorage := deps.NewCloudStorage(objectStoreConfig, clock)
	imagesCloudStorageService := &service.ImagesCloudStorageService{
		Storage: imagesCloudStorageServiceStorage,
	}
	factory := rootProvider.LoggerFactory
	getHandlerLogger := handler.NewGetHandlerLogger(factory)
	environmentConfig := &rootProvider.EnvironmentConfig
	imagesCDNHost := environmentConfig.ImagesCDNHost
	request := p.Request
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps2.ProvideHTTPHost(request, trustProxy)
	httpProto := deps2.ProvideHTTPProto(request, trustProxy)
	daemon := rootProvider.VipsDaemon
	getHandler := &handler.GetHandler{
		DirectorMaker: imagesCloudStorageService,
		Logger:        getHandlerLogger,
		ImagesCDNHost: imagesCDNHost,
		HTTPHost:      httpHost,
		HTTPProto:     httpProto,
		VipsDaemon:    daemon,
	}
	return getHandler
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)

func newPostHandler(p *deps.RequestProvider) http.Handler {
	appProvider := p.AppProvider
	rootProvider := appProvider.RootProvider
	factory := rootProvider.LoggerFactory
	postHandlerLogger := handler.NewPostHandlerLogger(factory)
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	objectStoreConfig := rootProvider.ObjectStoreConfig
	clockClock := _wireSystemClockValue
	imagesCloudStorageServiceStorage := deps.NewCloudStorage(objectStoreConfig, clockClock)
	imagesCloudStorageService := &service.ImagesCloudStorageService{
		Storage: imagesCloudStorageServiceStorage,
	}
	config := appProvider.Config
	secretConfig := config.SecretConfig
	imagesKeyMaterials := deps2.ProvideImagesKeyMaterials(secretConfig)
	request := p.Request
	environmentConfig := &rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps2.ProvideHTTPHost(request, trustProxy)
	provider := &presign.Provider{
		Secret: imagesKeyMaterials,
		Clock:  clockClock,
		Host:   httpHost,
	}
	pool := rootProvider.DatabasePool
	databaseEnvironmentConfig := environmentConfig.DatabaseConfig
	databaseCredentials := deps2.ProvideDatabaseCredentials(secretConfig)
	handle := appdb.NewHandle(pool, databaseEnvironmentConfig, databaseCredentials, factory)
	appConfig := config.AppConfig
	appID := appConfig.ID
	sqlBuilderApp := appdb.NewSQLBuilderApp(databaseCredentials, appID)
	sqlExecutor := appdb.NewSQLExecutor(handle)
	store := &images.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	postHandler := &handler.PostHandler{
		Logger:                         postHandlerLogger,
		JSON:                           jsonResponseWriter,
		PostHandlerCloudStorageService: imagesCloudStorageService,
		PresignProvider:                provider,
		Database:                       handle,
		ImagesStore:                    store,
		Clock:                          clockClock,
	}
	return postHandler
}
