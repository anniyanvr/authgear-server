// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package portal

import (
	"github.com/authgear/authgear-server/pkg/lib/admin/authz"
	"github.com/authgear/authgear-server/pkg/lib/analytic"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/config/plan"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/infra/middleware"
	"github.com/authgear/authgear-server/pkg/lib/tester"
	"github.com/authgear/authgear-server/pkg/lib/tutorial"
	"github.com/authgear/authgear-server/pkg/lib/usage"
	"github.com/authgear/authgear-server/pkg/portal/appresource/factory"
	"github.com/authgear/authgear-server/pkg/portal/appsecret"
	"github.com/authgear/authgear-server/pkg/portal/deps"
	"github.com/authgear/authgear-server/pkg/portal/endpoint"
	"github.com/authgear/authgear-server/pkg/portal/graphql"
	plan2 "github.com/authgear/authgear-server/pkg/portal/lib/plan"
	"github.com/authgear/authgear-server/pkg/portal/libstripe"
	"github.com/authgear/authgear-server/pkg/portal/loader"
	"github.com/authgear/authgear-server/pkg/portal/service"
	"github.com/authgear/authgear-server/pkg/portal/session"
	"github.com/authgear/authgear-server/pkg/portal/smtp"
	"github.com/authgear/authgear-server/pkg/portal/transport"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/intl"
	"github.com/authgear/authgear-server/pkg/util/template"
	"net/http"
)

import (
	_ "github.com/authgear/authgear-server/pkg/auth"
)

// Injectors from wire.go:

func newPanicMiddleware(p *deps.RequestProvider) httproute.Middleware {
	rootProvider := p.RootProvider
	factory := rootProvider.LoggerFactory
	panicMiddlewareLogger := middleware.NewPanicMiddlewareLogger(factory)
	panicMiddleware := &middleware.PanicMiddleware{
		Logger: panicMiddlewareLogger,
	}
	return panicMiddleware
}

func newBodyLimitMiddleware(p *deps.RequestProvider) httproute.Middleware {
	bodyLimitMiddleware := &middleware.BodyLimitMiddleware{}
	return bodyLimitMiddleware
}

func newSentryMiddleware(p *deps.RequestProvider) httproute.Middleware {
	rootProvider := p.RootProvider
	hub := rootProvider.SentryHub
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	sentryMiddleware := &middleware.SentryMiddleware{
		SentryHub:  hub,
		TrustProxy: trustProxy,
	}
	return sentryMiddleware
}

func newSessionInfoMiddleware(p *deps.RequestProvider) httproute.Middleware {
	sessionInfoMiddleware := &session.SessionInfoMiddleware{}
	return sessionInfoMiddleware
}

func newSessionRequiredMiddleware(p *deps.RequestProvider) httproute.Middleware {
	sessionRequiredMiddleware := &session.SessionRequiredMiddleware{}
	return sessionRequiredMiddleware
}

func newGraphQLHandler(p *deps.RequestProvider) http.Handler {
	request := p.Request
	rootProvider := p.RootProvider
	logFactory := rootProvider.LoggerFactory
	logger := graphql.NewLogger(logFactory)
	pool := rootProvider.Database
	environmentConfig := rootProvider.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	trustProxy := environmentConfig.TrustProxy
	authgearConfig := rootProvider.AuthgearConfig
	adminAPIConfig := rootProvider.AdminAPIConfig
	controller := rootProvider.ConfigSourceController
	configSource := deps.ProvideConfigSource(controller)
	clock := _wireSystemClockValue
	adder := &authz.Adder{
		Clock: clock,
	}
	appHostSuffixes := environmentConfig.AppHostSuffixes
	appConfig := rootProvider.AppConfig
	configServiceLogger := service.NewConfigServiceLogger(logFactory)
	domainImplementationType := rootProvider.DomainImplementation
	kubernetesConfig := rootProvider.KubernetesConfig
	kubernetesLogger := service.NewKubernetesLogger(logFactory)
	kubernetes := &service.Kubernetes{
		KubernetesConfig: kubernetesConfig,
		AppConfig:        appConfig,
		Logger:           kubernetesLogger,
	}
	configService := &service.ConfigService{
		Logger:               configServiceLogger,
		AppConfig:            appConfig,
		Controller:           controller,
		ConfigSource:         configSource,
		DomainImplementation: domainImplementationType,
		Kubernetes:           kubernetes,
	}
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(handle)
	domainService := &service.DomainService{
		Clock:          clock,
		DomainConfig:   configService,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		GlobalDatabase: handle,
	}
	defaultDomainService := &service.DefaultDomainService{
		AppHostSuffixes: appHostSuffixes,
		AppConfig:       appConfig,
		Domains:         domainService,
	}
	adminAPIService := &service.AdminAPIService{
		AuthgearConfig: authgearConfig,
		AdminAPIConfig: adminAPIConfig,
		ConfigSource:   configSource,
		AuthzAdder:     adder,
		DefaultDomains: defaultDomainService,
	}
	appServiceLogger := service.NewAppServiceLogger(logFactory)
	httpClient := service.NewHTTPClient()
	mailConfig := rootProvider.MailConfig
	smtpLogger := smtp.NewLogger(logFactory)
	devMode := environmentConfig.DevMode
	mailLogger := mail.NewLogger(logFactory)
	smtpConfig := rootProvider.SMTPConfig
	smtpServerCredentials := deps.ProvideSMTPServerCredentials(smtpConfig)
	dialer := mail.NewGomailDialer(smtpServerCredentials)
	sender := &mail.Sender{
		Logger:       mailLogger,
		GomailDialer: dialer,
	}
	smtpService := &smtp.Service{
		Logger:     smtpLogger,
		DevMode:    devMode,
		MailSender: sender,
	}
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	manager := rootProvider.Resources
	defaultLanguageTag := _wireDefaultLanguageTagValue
	supportedLanguageTags := _wireSupportedLanguageTagsValue
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	collaboratorService := &service.CollaboratorService{
		Clock:          clock,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		HTTPClient:     httpClient,
		GlobalDatabase: handle,
		MailConfig:     mailConfig,
		SMTPService:    smtpService,
		Endpoints:      endpointsProvider,
		TemplateEngine: engine,
		AdminAPI:       adminAPIService,
		AppConfigs:     configService,
	}
	authzService := &service.AuthzService{
		Configs:       configService,
		Collaborators: collaboratorService,
	}
	managerFactoryLogger := factory.NewManagerFactoryLogger(logFactory)
	appBaseResources := deps.ProvideAppBaseResources(rootProvider)
	storeImpl := &tutorial.StoreImpl{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	tutorialService := &tutorial.Service{
		GlobalDatabase: handle,
		Store:          storeImpl,
	}
	denoEndpoint := environmentConfig.DenoEndpoint
	hookLogger := hook.NewLogger(logFactory)
	denoClientImpl := ProvideDenoClient(denoEndpoint, hookLogger)
	managerFactory := &factory.ManagerFactory{
		Logger:            managerFactoryLogger,
		AppBaseResources:  appBaseResources,
		Tutorials:         tutorialService,
		DenoClient:        denoClientImpl,
		Clock:             clock,
		EnvironmentConfig: environmentConfig,
		DomainService:     domainService,
	}
	store := &plan.Store{
		Clock:       clock,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	planService := &plan2.Service{
		GlobalDatabase: handle,
		PlanStore:      store,
		AppConfig:      appConfig,
	}
	globalredisHandle := rootProvider.GlobalRedisHandle
	appSecretVisitTokenStoreImpl := &appsecret.AppSecretVisitTokenStoreImpl{
		Redis: globalredisHandle,
	}
	testerStore := &tester.TesterStore{
		Redis: globalredisHandle,
	}
	samlEnvironmentConfig := environmentConfig.SAML
	configsourceStore := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	appService := &service.AppService{
		Logger:                   appServiceLogger,
		SQLBuilder:               sqlBuilder,
		SQLExecutor:              sqlExecutor,
		GlobalDatabase:           handle,
		AppConfig:                appConfig,
		AppConfigs:               configService,
		AppAuthz:                 authzService,
		DefaultDomains:           defaultDomainService,
		Resources:                manager,
		AppResMgrFactory:         managerFactory,
		Plan:                     planService,
		Clock:                    clock,
		AppSecretVisitTokenStore: appSecretVisitTokenStoreImpl,
		AppTesterTokenStore:      testerStore,
		SAMLEnvironmentConfig:    samlEnvironmentConfig,
		ConfigSourceStore:        configsourceStore,
	}
	loaderHTTPClient := loader.NewHTTPClient()
	userLoader := loader.NewUserLoader(adminAPIService, appService, collaboratorService, loaderHTTPClient)
	appLoader := loader.NewAppLoader(appService, authzService)
	domainLoader := loader.NewDomainLoader(domainService, authzService)
	collaboratorLoader := loader.NewCollaboratorLoader(collaboratorService, authzService)
	collaboratorInvitationLoader := loader.NewCollaboratorInvitationLoader(collaboratorService, authzService)
	auditDatabaseCredentials := deps.ProvideAuditDatabaseCredentials(environmentConfig)
	readHandle := auditdb.NewReadHandle(pool, databaseEnvironmentConfig, auditDatabaseCredentials, logFactory)
	auditdbSQLBuilder := auditdb.NewSQLBuilder(auditDatabaseCredentials)
	readSQLExecutor := auditdb.NewReadSQLExecutor(readHandle)
	auditDBReadStore := &analytic.AuditDBReadStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: readSQLExecutor,
	}
	analyticConfig := rootProvider.AnalyticConfig
	chartService := &analytic.ChartService{
		Database:       readHandle,
		AuditStore:     auditDBReadStore,
		Clock:          clock,
		AnalyticConfig: analyticConfig,
	}
	stripeConfig := rootProvider.StripeConfig
	libstripeLogger := libstripe.NewLogger(logFactory)
	api := libstripe.NewClientAPI(stripeConfig, libstripeLogger)
	stripeCache := libstripe.NewStripeCache()
	libstripeService := &libstripe.Service{
		ClientAPI:         api,
		Logger:            libstripeLogger,
		Plans:             planService,
		GlobalRedisHandle: globalredisHandle,
		Cache:             stripeCache,
		Clock:             clock,
		StripeConfig:      stripeConfig,
		Endpoints:         endpointsProvider,
	}
	globalDBStore := &usage.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	subscriptionService := &service.SubscriptionService{
		SQLBuilder:        sqlBuilder,
		SQLExecutor:       sqlExecutor,
		GlobalDatabase:    handle,
		ConfigSourceStore: configsourceStore,
		PlanStore:         store,
		UsageStore:        globalDBStore,
		Clock:             clock,
		AppConfig:         appConfig,
	}
	usageService := &service.UsageService{
		GlobalDatabase: handle,
		UsageStore:     globalDBStore,
	}
	remoteIP := deps.ProvideRemoteIP(request, trustProxy)
	userAgentString := deps.ProvideUserAgentString(request)
	writeHandle := auditdb.NewWriteHandle(pool, databaseEnvironmentConfig, auditDatabaseCredentials, logFactory)
	auditService := &service.AuditService{
		RemoteIP:          remoteIP,
		UserAgentString:   userAgentString,
		Request:           request,
		Apps:              appService,
		Authgear:          authgearConfig,
		DenoEndpoint:      denoEndpoint,
		GlobalSQLBuilder:  sqlBuilder,
		GlobalSQLExecutor: sqlExecutor,
		GlobalDatabase:    handle,
		AuditDatabase:     writeHandle,
		Clock:             clock,
		LoggerFactory:     logFactory,
	}
	onboardService := &service.OnboardService{
		HTTPClient:     httpClient,
		AuthgearConfig: authgearConfig,
		AdminAPI:       adminAPIService,
	}
	context := &graphql.Context{
		Request:                 request,
		GQLLogger:               logger,
		GlobalDatabase:          handle,
		TrustProxy:              trustProxy,
		Users:                   userLoader,
		Apps:                    appLoader,
		Domains:                 domainLoader,
		Collaborators:           collaboratorLoader,
		CollaboratorInvitations: collaboratorInvitationLoader,
		AuthzService:            authzService,
		AppService:              appService,
		DomainService:           domainService,
		CollaboratorService:     collaboratorService,
		SMTPService:             smtpService,
		AppResMgrFactory:        managerFactory,
		AnalyticChartService:    chartService,
		TutorialService:         tutorialService,
		StripeService:           libstripeService,
		SubscriptionService:     subscriptionService,
		UsageService:            usageService,
		DenoService:             denoClientImpl,
		AuditService:            auditService,
		OnboardService:          onboardService,
	}
	graphQLHandler := &transport.GraphQLHandler{
		GraphQLContext: context,
	}
	return graphQLHandler
}

var (
	_wireSystemClockValue           = clock.NewSystemClock()
	_wireDefaultLanguageTagValue    = template.DefaultLanguageTag(intl.BuiltinBaseLanguage)
	_wireSupportedLanguageTagsValue = template.SupportedLanguageTags([]string{intl.BuiltinBaseLanguage})
)

func newSystemConfigHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	authgearConfig := rootProvider.AuthgearConfig
	appConfig := rootProvider.AppConfig
	searchConfig := rootProvider.SearchConfig
	auditLogConfig := rootProvider.AuditLogConfig
	analyticConfig := rootProvider.AnalyticConfig
	googleTagManagerConfig := rootProvider.GoogleTagManagerConfig
	portalFrontendSentryConfig := rootProvider.PortalFrontendSentryConfig
	environmentConfig := rootProvider.EnvironmentConfig
	globalUIImplementation := environmentConfig.UIImplementation
	globalUISettingsImplementation := environmentConfig.UISettingsImplementation
	manager := rootProvider.Resources
	systemConfigProvider := &service.SystemConfigProvider{
		AuthgearConfig:                 authgearConfig,
		AppConfig:                      appConfig,
		SearchConfig:                   searchConfig,
		AuditLogConfig:                 auditLogConfig,
		AnalyticConfig:                 analyticConfig,
		GTMConfig:                      googleTagManagerConfig,
		FrontendSentryConfig:           portalFrontendSentryConfig,
		GlobalUIImplementation:         globalUIImplementation,
		GlobalUISettingsImplementation: globalUISettingsImplementation,
		Resources:                      manager,
	}
	filesystemCache := rootProvider.FilesystemCache
	systemConfigHandler := &transport.SystemConfigHandler{
		SystemConfig:    systemConfigProvider,
		FilesystemCache: filesystemCache,
	}
	return systemConfigHandler
}

func newAdminAPIHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	pool := rootProvider.Database
	environmentConfig := rootProvider.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	logFactory := rootProvider.LoggerFactory
	handle := globaldb.NewHandle(pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	configServiceLogger := service.NewConfigServiceLogger(logFactory)
	appConfig := rootProvider.AppConfig
	controller := rootProvider.ConfigSourceController
	configSource := deps.ProvideConfigSource(controller)
	domainImplementationType := rootProvider.DomainImplementation
	kubernetesConfig := rootProvider.KubernetesConfig
	kubernetesLogger := service.NewKubernetesLogger(logFactory)
	kubernetes := &service.Kubernetes{
		KubernetesConfig: kubernetesConfig,
		AppConfig:        appConfig,
		Logger:           kubernetesLogger,
	}
	configService := &service.ConfigService{
		Logger:               configServiceLogger,
		AppConfig:            appConfig,
		Controller:           controller,
		ConfigSource:         configSource,
		DomainImplementation: domainImplementationType,
		Kubernetes:           kubernetes,
	}
	clockClock := _wireSystemClockValue
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(handle)
	httpClient := service.NewHTTPClient()
	mailConfig := rootProvider.MailConfig
	logger := smtp.NewLogger(logFactory)
	devMode := environmentConfig.DevMode
	mailLogger := mail.NewLogger(logFactory)
	smtpConfig := rootProvider.SMTPConfig
	smtpServerCredentials := deps.ProvideSMTPServerCredentials(smtpConfig)
	dialer := mail.NewGomailDialer(smtpServerCredentials)
	sender := &mail.Sender{
		Logger:       mailLogger,
		GomailDialer: dialer,
	}
	smtpService := &smtp.Service{
		Logger:     logger,
		DevMode:    devMode,
		MailSender: sender,
	}
	request := p.Request
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	manager := rootProvider.Resources
	defaultLanguageTag := _wireDefaultLanguageTagValue
	supportedLanguageTags := _wireSupportedLanguageTagsValue
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	authgearConfig := rootProvider.AuthgearConfig
	adminAPIConfig := rootProvider.AdminAPIConfig
	adder := &authz.Adder{
		Clock: clockClock,
	}
	appHostSuffixes := environmentConfig.AppHostSuffixes
	domainService := &service.DomainService{
		Clock:          clockClock,
		DomainConfig:   configService,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		GlobalDatabase: handle,
	}
	defaultDomainService := &service.DefaultDomainService{
		AppHostSuffixes: appHostSuffixes,
		AppConfig:       appConfig,
		Domains:         domainService,
	}
	adminAPIService := &service.AdminAPIService{
		AuthgearConfig: authgearConfig,
		AdminAPIConfig: adminAPIConfig,
		ConfigSource:   configSource,
		AuthzAdder:     adder,
		DefaultDomains: defaultDomainService,
	}
	collaboratorService := &service.CollaboratorService{
		Clock:          clockClock,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		HTTPClient:     httpClient,
		GlobalDatabase: handle,
		MailConfig:     mailConfig,
		SMTPService:    smtpService,
		Endpoints:      endpointsProvider,
		TemplateEngine: engine,
		AdminAPI:       adminAPIService,
		AppConfigs:     configService,
	}
	authzService := &service.AuthzService{
		Configs:       configService,
		Collaborators: collaboratorService,
	}
	adminAPILogger := transport.NewAdminAPILogger(logFactory)
	adminAPIHandler := &transport.AdminAPIHandler{
		Database: handle,
		Authz:    authzService,
		AdminAPI: adminAPIService,
		Logger:   adminAPILogger,
	}
	return adminAPIHandler
}

func newStaticAssetsHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	manager := rootProvider.Resources
	staticAssetsHandler := &transport.StaticAssetsHandler{
		Resources: manager,
	}
	return staticAssetsHandler
}

func newStripeWebhookHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	stripeConfig := rootProvider.StripeConfig
	logFactory := rootProvider.LoggerFactory
	logger := libstripe.NewLogger(logFactory)
	api := libstripe.NewClientAPI(stripeConfig, logger)
	pool := rootProvider.Database
	environmentConfig := rootProvider.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	clockClock := _wireSystemClockValue
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(handle)
	store := &plan.Store{
		Clock:       clockClock,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	appConfig := rootProvider.AppConfig
	planService := &plan2.Service{
		GlobalDatabase: handle,
		PlanStore:      store,
		AppConfig:      appConfig,
	}
	globalredisHandle := rootProvider.GlobalRedisHandle
	stripeCache := libstripe.NewStripeCache()
	request := p.Request
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	libstripeService := &libstripe.Service{
		ClientAPI:         api,
		Logger:            logger,
		Plans:             planService,
		GlobalRedisHandle: globalredisHandle,
		Cache:             stripeCache,
		Clock:             clockClock,
		StripeConfig:      stripeConfig,
		Endpoints:         endpointsProvider,
	}
	stripeWebhookLogger := transport.NewStripeWebhookLogger(logFactory)
	configsourceStore := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	globalDBStore := &usage.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	subscriptionService := &service.SubscriptionService{
		SQLBuilder:        sqlBuilder,
		SQLExecutor:       sqlExecutor,
		GlobalDatabase:    handle,
		ConfigSourceStore: configsourceStore,
		PlanStore:         store,
		UsageStore:        globalDBStore,
		Clock:             clockClock,
		AppConfig:         appConfig,
	}
	stripeWebhookHandler := &transport.StripeWebhookHandler{
		StripeService: libstripeService,
		Logger:        stripeWebhookLogger,
		Subscriptions: subscriptionService,
		Database:      handle,
	}
	return stripeWebhookHandler
}

func newOsanoHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	osanoConfig := rootProvider.OsanoConfig
	osanoHandler := &transport.OsanoHandler{
		OsanoConfig: osanoConfig,
	}
	return osanoHandler
}
