// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package admin

import (
	"context"
	facade2 "github.com/authgear/authgear-server/pkg/admin/facade"
	"github.com/authgear/authgear-server/pkg/admin/graphql"
	"github.com/authgear/authgear-server/pkg/admin/loader"
	service3 "github.com/authgear/authgear-server/pkg/admin/service"
	"github.com/authgear/authgear-server/pkg/admin/transport"
	"github.com/authgear/authgear-server/pkg/lib/admin/authz"
	"github.com/authgear/authgear-server/pkg/lib/audit"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticationinfo"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/oob"
	passkey3 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/passkey"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/password"
	service2 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/totp"
	"github.com/authgear/authgear-server/pkg/lib/authn/challenge"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/anonymous"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/biometric"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/passkey"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/siwe"
	"github.com/authgear/authgear-server/pkg/lib/authn/mfa"
	"github.com/authgear/authgear-server/pkg/lib/authn/otp"
	"github.com/authgear/authgear-server/pkg/lib/authn/sso"
	stdattrs2 "github.com/authgear/authgear-server/pkg/lib/authn/stdattrs"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/elasticsearch"
	"github.com/authgear/authgear-server/pkg/lib/endpoints"
	"github.com/authgear/authgear-server/pkg/lib/event"
	"github.com/authgear/authgear-server/pkg/lib/facade"
	"github.com/authgear/authgear-server/pkg/lib/feature/customattrs"
	"github.com/authgear/authgear-server/pkg/lib/feature/forgotpassword"
	passkey2 "github.com/authgear/authgear-server/pkg/lib/feature/passkey"
	siwe2 "github.com/authgear/authgear-server/pkg/lib/feature/siwe"
	"github.com/authgear/authgear-server/pkg/lib/feature/stdattrs"
	"github.com/authgear/authgear-server/pkg/lib/feature/verification"
	"github.com/authgear/authgear-server/pkg/lib/feature/web3"
	"github.com/authgear/authgear-server/pkg/lib/healthz"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/middleware"
	"github.com/authgear/authgear-server/pkg/lib/infra/whatsapp"
	"github.com/authgear/authgear-server/pkg/lib/interaction"
	"github.com/authgear/authgear-server/pkg/lib/lockout"
	"github.com/authgear/authgear-server/pkg/lib/messaging"
	"github.com/authgear/authgear-server/pkg/lib/nonce"
	oauth2 "github.com/authgear/authgear-server/pkg/lib/oauth"
	"github.com/authgear/authgear-server/pkg/lib/oauth/handler"
	"github.com/authgear/authgear-server/pkg/lib/oauth/oidc"
	"github.com/authgear/authgear-server/pkg/lib/oauth/pq"
	"github.com/authgear/authgear-server/pkg/lib/oauth/redis"
	"github.com/authgear/authgear-server/pkg/lib/presign"
	"github.com/authgear/authgear-server/pkg/lib/ratelimit"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/lib/session/access"
	"github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	"github.com/authgear/authgear-server/pkg/lib/sessionlisting"
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/lib/usage"
	"github.com/authgear/authgear-server/pkg/lib/web"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/httputil"
	"github.com/authgear/authgear-server/pkg/util/rand"
	"github.com/authgear/authgear-server/pkg/util/template"
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

func newHealthzHandler(p *deps.RootProvider, w http.ResponseWriter, r *http.Request, ctx context.Context) http.Handler {
	pool := p.DatabasePool
	environmentConfig := p.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	factory := p.LoggerFactory
	handle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlExecutor := globaldb.NewSQLExecutor(ctx, handle)
	handlerLogger := healthz.NewHandlerLogger(factory)
	handler := &healthz.Handler{
		Context:        ctx,
		GlobalDatabase: handle,
		GlobalExecutor: sqlExecutor,
		Logger:         handlerLogger,
	}
	return handler
}

func newSentryMiddleware(p *deps.RootProvider) httproute.Middleware {
	hub := p.SentryHub
	environmentConfig := p.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	sentryMiddleware := &middleware.SentryMiddleware{
		SentryHub:  hub,
		TrustProxy: trustProxy,
	}
	return sentryMiddleware
}

func newBodyLimitMiddleware(p *deps.RootProvider) httproute.Middleware {
	bodyLimitMiddleware := &middleware.BodyLimitMiddleware{}
	return bodyLimitMiddleware
}

func newAuthorizationMiddleware(p *deps.RequestProvider, auth config.AdminAPIAuth) httproute.Middleware {
	appProvider := p.AppProvider
	factory := appProvider.LoggerFactory
	logger := authz.NewLogger(factory)
	appContext := appProvider.AppContext
	configConfig := appContext.Config
	appConfig := configConfig.AppConfig
	appID := appConfig.ID
	secretConfig := configConfig.SecretConfig
	adminAPIAuthKey := deps.ProvideAdminAPIAuthKeyMaterials(secretConfig)
	clock := _wireSystemClockValue
	authzMiddleware := &authz.Middleware{
		Logger:  logger,
		Auth:    auth,
		AppID:   appID,
		AuthKey: adminAPIAuthKey,
		Clock:   clock,
	}
	return authzMiddleware
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)

func newUIParamMiddleware(p *deps.RequestProvider) httproute.Middleware {
	uiParamMiddleware := &transport.UIParamMiddleware{}
	return uiParamMiddleware
}

func newGraphQLHandler(p *deps.RequestProvider) http.Handler {
	appProvider := p.AppProvider
	factory := appProvider.LoggerFactory
	logger := graphql.NewLogger(factory)
	appContext := appProvider.AppContext
	configConfig := appContext.Config
	appConfig := configConfig.AppConfig
	oAuthConfig := appConfig.OAuth
	featureConfig := configConfig.FeatureConfig
	adminAPIFeatureConfig := featureConfig.AdminAPI
	secretConfig := configConfig.SecretConfig
	databaseCredentials := deps.ProvideDatabaseCredentials(secretConfig)
	appID := appConfig.ID
	sqlBuilderApp := appdb.NewSQLBuilderApp(databaseCredentials, appID)
	request := p.Request
	contextContext := deps.ProvideRequestContext(request)
	handle := appProvider.AppDatabase
	sqlExecutor := appdb.NewSQLExecutor(contextContext, handle)
	clockClock := _wireSystemClockValue
	store := &user.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	rawQueries := &user.RawQueries{
		Store: store,
	}
	authenticationConfig := appConfig.Authentication
	identityConfig := appConfig.Identity
	identityFeatureConfig := featureConfig.Identity
	serviceStore := &service.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	loginidStore := &loginid.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	loginIDConfig := identityConfig.LoginID
	manager := appContext.Resources
	typeCheckerFactory := &loginid.TypeCheckerFactory{
		Config:    loginIDConfig,
		Resources: manager,
	}
	checker := &loginid.Checker{
		Config:             loginIDConfig,
		TypeCheckerFactory: typeCheckerFactory,
	}
	normalizerFactory := &loginid.NormalizerFactory{
		Config: loginIDConfig,
	}
	provider := &loginid.Provider{
		Store:             loginidStore,
		Config:            loginIDConfig,
		Checker:           checker,
		NormalizerFactory: normalizerFactory,
		Clock:             clockClock,
	}
	oauthStore := &oauth.Store{
		SQLBuilder:     sqlBuilderApp,
		SQLExecutor:    sqlExecutor,
		IdentityConfig: identityConfig,
	}
	oauthProvider := &oauth.Provider{
		Store:          oauthStore,
		Clock:          clockClock,
		IdentityConfig: identityConfig,
	}
	anonymousStore := &anonymous.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	anonymousProvider := &anonymous.Provider{
		Store: anonymousStore,
		Clock: clockClock,
	}
	biometricStore := &biometric.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	biometricProvider := &biometric.Provider{
		Store: biometricStore,
		Clock: clockClock,
	}
	passkeyStore := &passkey.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	appredisHandle := appProvider.Redis
	store2 := &passkey2.Store{
		Context: contextContext,
		Redis:   appredisHandle,
		AppID:   appID,
	}
	rootProvider := appProvider.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	defaultLanguageTag := deps.ProvideDefaultLanguageTag(configConfig)
	supportedLanguageTags := deps.ProvideSupportedLanguageTags(configConfig)
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	httpConfig := appConfig.HTTP
	localizationConfig := appConfig.Localization
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	webAppCDNHost := environmentConfig.WebAppCDNHost
	globalEmbeddedResourceManager := rootProvider.EmbeddedResources
	staticAssetResolver := &web.StaticAssetResolver{
		Context:           contextContext,
		Config:            httpConfig,
		Localization:      localizationConfig,
		HTTPProto:         httpProto,
		WebAppCDNHost:     webAppCDNHost,
		Resources:         manager,
		EmbeddedResources: globalEmbeddedResourceManager,
	}
	translationService := &translation.Service{
		Context:        contextContext,
		TemplateEngine: engine,
		StaticAssets:   staticAssetResolver,
	}
	configService := &passkey2.ConfigService{
		Request:            request,
		TrustProxy:         trustProxy,
		TranslationService: translationService,
	}
	passkeyService := &passkey2.Service{
		Store:         store2,
		ConfigService: configService,
	}
	passkeyProvider := &passkey.Provider{
		Store:   passkeyStore,
		Clock:   clockClock,
		Passkey: passkeyService,
	}
	siweStore := &siwe.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	remoteIP := deps.ProvideRemoteIP(request, trustProxy)
	web3Config := appConfig.Web3
	storeRedis := &siwe2.StoreRedis{
		Context: contextContext,
		Redis:   appredisHandle,
		AppID:   appID,
		Clock:   clockClock,
	}
	ratelimitLogger := ratelimit.NewLogger(factory)
	storageRedis := &ratelimit.StorageRedis{
		AppID: appID,
		Redis: appredisHandle,
	}
	rateLimitsFeatureConfig := featureConfig.RateLimits
	limiter := &ratelimit.Limiter{
		Logger:  ratelimitLogger,
		Storage: storageRedis,
		Config:  rateLimitsFeatureConfig,
	}
	siweLogger := siwe2.NewLogger(factory)
	siweService := &siwe2.Service{
		RemoteIP:             remoteIP,
		HTTPConfig:           httpConfig,
		Web3Config:           web3Config,
		AuthenticationConfig: authenticationConfig,
		Clock:                clockClock,
		NonceStore:           storeRedis,
		RateLimiter:          limiter,
		Logger:               siweLogger,
	}
	siweProvider := &siwe.Provider{
		Store: siweStore,
		Clock: clockClock,
		SIWE:  siweService,
	}
	serviceService := &service.Service{
		Authentication:        authenticationConfig,
		Identity:              identityConfig,
		IdentityFeatureConfig: identityFeatureConfig,
		Store:                 serviceStore,
		LoginID:               provider,
		OAuth:                 oauthProvider,
		Anonymous:             anonymousProvider,
		Biometric:             biometricProvider,
		Passkey:               passkeyProvider,
		SIWE:                  siweProvider,
	}
	store3 := &service2.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	passwordStore := &password.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	authenticatorConfig := appConfig.Authenticator
	authenticatorPasswordConfig := authenticatorConfig.Password
	passwordLogger := password.NewLogger(factory)
	historyStore := &password.HistoryStore{
		Clock:       clockClock,
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	authenticatorFeatureConfig := featureConfig.Authenticator
	passwordChecker := password.ProvideChecker(authenticatorPasswordConfig, authenticatorFeatureConfig, historyStore)
	housekeeperLogger := password.NewHousekeeperLogger(factory)
	housekeeper := &password.Housekeeper{
		Store:  historyStore,
		Logger: housekeeperLogger,
		Config: authenticatorPasswordConfig,
	}
	passwordProvider := &password.Provider{
		Store:           passwordStore,
		Config:          authenticatorPasswordConfig,
		Clock:           clockClock,
		Logger:          passwordLogger,
		PasswordHistory: historyStore,
		PasswordChecker: passwordChecker,
		Housekeeper:     housekeeper,
	}
	store4 := &passkey3.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	provider2 := &passkey3.Provider{
		Store:   store4,
		Clock:   clockClock,
		Passkey: passkeyService,
	}
	totpStore := &totp.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	authenticatorTOTPConfig := authenticatorConfig.TOTP
	totpProvider := &totp.Provider{
		Store:  totpStore,
		Config: authenticatorTOTPConfig,
		Clock:  clockClock,
	}
	oobStore := &oob.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	oobProvider := &oob.Provider{
		Store: oobStore,
		Clock: clockClock,
	}
	testModeFeatureConfig := featureConfig.TestMode
	codeStoreRedis := &otp.CodeStoreRedis{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	lookupStoreRedis := &otp.LookupStoreRedis{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	attemptTrackerRedis := &otp.AttemptTrackerRedis{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	otpLogger := otp.NewLogger(factory)
	otpService := &otp.Service{
		Clock:                 clockClock,
		AppID:                 appID,
		TestModeFeatureConfig: testModeFeatureConfig,
		RemoteIP:              remoteIP,
		CodeStore:             codeStoreRedis,
		LookupStore:           lookupStoreRedis,
		AttemptTracker:        attemptTrackerRedis,
		Logger:                otpLogger,
		RateLimiter:           limiter,
	}
	rateLimits := service2.RateLimits{
		IP:          remoteIP,
		Config:      authenticationConfig,
		RateLimiter: limiter,
	}
	authenticationLockoutConfig := authenticationConfig.Lockout
	lockoutLogger := lockout.NewLogger(factory)
	lockoutStorageRedis := &lockout.StorageRedis{
		AppID: appID,
		Redis: appredisHandle,
	}
	lockoutService := &lockout.Service{
		Logger:  lockoutLogger,
		Storage: lockoutStorageRedis,
	}
	serviceLockout := service2.Lockout{
		Config:   authenticationLockoutConfig,
		RemoteIP: remoteIP,
		provider: lockoutService,
	}
	service4 := &service2.Service{
		Store:          store3,
		Config:         appConfig,
		Password:       passwordProvider,
		Passkey:        provider2,
		TOTP:           totpProvider,
		OOBOTP:         oobProvider,
		OTPCodeService: otpService,
		RateLimits:     rateLimits,
		Lockout:        serviceLockout,
	}
	verificationConfig := appConfig.Verification
	userProfileConfig := appConfig.UserProfile
	storePQ := &verification.StorePQ{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	verificationService := &verification.Service{
		Config:            verificationConfig,
		UserProfileConfig: userProfileConfig,
		Clock:             clockClock,
		ClaimStore:        storePQ,
	}
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	imagesCDNHost := environmentConfig.ImagesCDNHost
	pictureTransformer := &stdattrs.PictureTransformer{
		HTTPProto:     httpProto,
		HTTPHost:      httpHost,
		ImagesCDNHost: imagesCDNHost,
	}
	serviceNoEvent := &stdattrs.ServiceNoEvent{
		UserProfileConfig: userProfileConfig,
		Identities:        serviceService,
		UserQueries:       rawQueries,
		UserStore:         store,
		ClaimStore:        storePQ,
		Transformer:       pictureTransformer,
	}
	customattrsServiceNoEvent := &customattrs.ServiceNoEvent{
		Config:      userProfileConfig,
		UserQueries: rawQueries,
		UserStore:   store,
	}
	nftIndexerAPIEndpoint := environmentConfig.NFTIndexerAPIEndpoint
	web3Service := &web3.Service{
		APIEndpoint: nftIndexerAPIEndpoint,
		Web3Config:  web3Config,
	}
	queries := &user.Queries{
		RawQueries:         rawQueries,
		Store:              store,
		Identities:         serviceService,
		Authenticators:     service4,
		Verification:       verificationService,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
		Web3:               web3Service,
	}
	userLoader := loader.NewUserLoader(queries)
	identityLoader := loader.NewIdentityLoader(serviceService)
	authenticatorLoader := loader.NewAuthenticatorLoader(service4)
	readHandle := appProvider.AuditReadDatabase
	auditDatabaseCredentials := deps.ProvideAuditDatabaseCredentials(secretConfig)
	auditdbSQLBuilderApp := auditdb.NewSQLBuilderApp(auditDatabaseCredentials, appID)
	readSQLExecutor := auditdb.NewReadSQLExecutor(contextContext, readHandle)
	readStore := &audit.ReadStore{
		SQLBuilder:  auditdbSQLBuilderApp,
		SQLExecutor: readSQLExecutor,
	}
	query := &audit.Query{
		Database: readHandle,
		Store:    readStore,
	}
	auditLogLoader := loader.NewAuditLogLoader(query, readHandle)
	elasticsearchCredentials := deps.ProvideElasticsearchCredentials(secretConfig)
	client := elasticsearch.NewClient(elasticsearchCredentials)
	queue := appProvider.TaskQueue
	elasticsearchService := &elasticsearch.Service{
		AppID:     appID,
		Client:    client,
		Users:     queries,
		OAuth:     oauthStore,
		LoginID:   loginidStore,
		TaskQueue: queue,
	}
	rawCommands := &user.RawCommands{
		Store: store,
		Clock: clockClock,
	}
	userAgentString := deps.ProvideUserAgentString(request)
	eventLogger := event.NewLogger(factory)
	sqlBuilder := appdb.NewSQLBuilder(databaseCredentials)
	storeImpl := event.NewStoreImpl(sqlBuilder, sqlExecutor)
	resolverImpl := &event.ResolverImpl{
		Users: queries,
	}
	hookLogger := hook.NewLogger(factory)
	hookConfig := appConfig.Hook
	webhookKeyMaterials := deps.ProvideWebhookKeyMaterials(secretConfig)
	webHookImpl := hook.WebHookImpl{
		Secret: webhookKeyMaterials,
	}
	syncHTTPClient := hook.NewSyncHTTPClient(hookConfig)
	asyncHTTPClient := hook.NewAsyncHTTPClient()
	eventWebHookImpl := &hook.EventWebHookImpl{
		WebHookImpl: webHookImpl,
		SyncHTTP:    syncHTTPClient,
		AsyncHTTP:   asyncHTTPClient,
	}
	denoHook := hook.DenoHook{
		Context:         contextContext,
		ResourceManager: manager,
	}
	denoEndpoint := environmentConfig.DenoEndpoint
	syncDenoClient := hook.NewSyncDenoClient(denoEndpoint, hookConfig, hookLogger)
	asyncDenoClient := hook.NewAsyncDenoClient(denoEndpoint, hookLogger)
	eventDenoHookImpl := &hook.EventDenoHookImpl{
		DenoHook:        denoHook,
		SyncDenoClient:  syncDenoClient,
		AsyncDenoClient: asyncDenoClient,
	}
	sink := &hook.Sink{
		Logger:             hookLogger,
		Config:             hookConfig,
		Clock:              clockClock,
		EventWebHook:       eventWebHookImpl,
		EventDenoHook:      eventDenoHookImpl,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
	}
	auditLogger := audit.NewLogger(factory)
	writeHandle := appProvider.AuditWriteDatabase
	writeSQLExecutor := auditdb.NewWriteSQLExecutor(contextContext, writeHandle)
	writeStore := &audit.WriteStore{
		SQLBuilder:  auditdbSQLBuilderApp,
		SQLExecutor: writeSQLExecutor,
	}
	auditSink := &audit.Sink{
		Logger:   auditLogger,
		Database: writeHandle,
		Store:    writeStore,
	}
	elasticsearchLogger := elasticsearch.NewLogger(factory)
	service5 := elasticsearch.Service{
		AppID:     appID,
		Client:    client,
		Users:     queries,
		OAuth:     oauthStore,
		LoginID:   loginidStore,
		TaskQueue: queue,
	}
	elasticsearchSink := &elasticsearch.Sink{
		Logger:   elasticsearchLogger,
		Service:  service5,
		Database: handle,
	}
	eventService := event.NewService(contextContext, remoteIP, userAgentString, eventLogger, handle, clockClock, localizationConfig, storeImpl, resolverImpl, sink, auditSink, elasticsearchSink)
	commands := &user.Commands{
		RawCommands:        rawCommands,
		RawQueries:         rawQueries,
		Events:             eventService,
		Verification:       verificationService,
		UserProfileConfig:  userProfileConfig,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
		Web3:               web3Service,
	}
	userProvider := &user.Provider{
		Commands: commands,
		Queries:  queries,
	}
	storeDeviceTokenRedis := &mfa.StoreDeviceTokenRedis{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	storeRecoveryCodePQ := &mfa.StoreRecoveryCodePQ{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	mfaLockout := mfa.Lockout{
		Config:   authenticationLockoutConfig,
		RemoteIP: remoteIP,
		provider: lockoutService,
	}
	mfaService := &mfa.Service{
		IP:            remoteIP,
		DeviceTokens:  storeDeviceTokenRedis,
		RecoveryCodes: storeRecoveryCodePQ,
		Clock:         clockClock,
		Config:        authenticationConfig,
		RateLimiter:   limiter,
		Lockout:       mfaLockout,
	}
	stdattrsService := &stdattrs.Service{
		UserProfileConfig: userProfileConfig,
		ServiceNoEvent:    serviceNoEvent,
		Identities:        serviceService,
		UserQueries:       rawQueries,
		UserStore:         store,
		Events:            eventService,
	}
	authorizationStore := &pq.AuthorizationStore{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	storeRedisLogger := idpsession.NewStoreRedisLogger(factory)
	idpsessionStoreRedis := &idpsession.StoreRedis{
		Redis:  appredisHandle,
		AppID:  appID,
		Clock:  clockClock,
		Logger: storeRedisLogger,
	}
	sessionConfig := appConfig.Session
	cookieManager := deps.NewCookieManager(request, trustProxy, httpConfig)
	cookieDef := session.NewSessionCookieDef(sessionConfig)
	idpsessionManager := &idpsession.Manager{
		Store:     idpsessionStoreRedis,
		Config:    sessionConfig,
		Cookies:   cookieManager,
		CookieDef: cookieDef,
	}
	redisLogger := redis.NewLogger(factory)
	redisStore := &redis.Store{
		Context:     contextContext,
		Redis:       appredisHandle,
		AppID:       appID,
		Logger:      redisLogger,
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	eventStoreRedis := &access.EventStoreRedis{
		Redis: appredisHandle,
		AppID: appID,
	}
	eventProvider := &access.EventProvider{
		Store: eventStoreRedis,
	}
	rand := _wireRandValue
	idpsessionProvider := &idpsession.Provider{
		Context:         contextContext,
		RemoteIP:        remoteIP,
		UserAgentString: userAgentString,
		AppID:           appID,
		Redis:           appredisHandle,
		Store:           idpsessionStoreRedis,
		AccessEvents:    eventProvider,
		TrustProxy:      trustProxy,
		Config:          sessionConfig,
		Clock:           clockClock,
		Random:          rand,
	}
	offlineGrantService := oauth2.OfflineGrantService{
		OAuthConfig: oAuthConfig,
		Clock:       clockClock,
		IDPSessions: idpsessionProvider,
	}
	sessionManager := &oauth2.SessionManager{
		Store:   redisStore,
		Config:  oAuthConfig,
		Service: offlineGrantService,
	}
	accountDeletionConfig := appConfig.AccountDeletion
	accountAnonymizationConfig := appConfig.AccountAnonymization
	coordinator := &facade.Coordinator{
		Events:                     eventService,
		Identities:                 serviceService,
		Authenticators:             service4,
		Verification:               verificationService,
		MFA:                        mfaService,
		UserCommands:               commands,
		UserQueries:                queries,
		StdAttrsService:            stdattrsService,
		PasswordHistory:            historyStore,
		OAuth:                      authorizationStore,
		IDPSessions:                idpsessionManager,
		OAuthSessions:              sessionManager,
		IdentityConfig:             identityConfig,
		AccountDeletionConfig:      accountDeletionConfig,
		AccountAnonymizationConfig: accountAnonymizationConfig,
		Clock:                      clockClock,
	}
	userFacade := &facade.UserFacade{
		UserProvider: userProvider,
		Coordinator:  coordinator,
	}
	interactionLogger := interaction.NewLogger(factory)
	identityFacade := facade.IdentityFacade{
		Coordinator: coordinator,
	}
	authenticatorFacade := facade.AuthenticatorFacade{
		Coordinator: coordinator,
	}
	anonymousStoreRedis := &anonymous.StoreRedis{
		Context: contextContext,
		Redis:   appredisHandle,
		AppID:   appID,
		Clock:   clockClock,
	}
	endpointsEndpoints := &endpoints.Endpoints{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	messagingLogger := messaging.NewLogger(factory)
	usageLogger := usage.NewLogger(factory)
	usageLimiter := &usage.Limiter{
		Logger: usageLogger,
		Clock:  clockClock,
		AppID:  appID,
		Redis:  appredisHandle,
	}
	messagingConfig := appConfig.Messaging
	messagingRateLimitsConfig := messagingConfig.RateLimits
	messagingFeatureConfig := featureConfig.Messaging
	rateLimitsEnvironmentConfig := &environmentConfig.RateLimits
	limits := messaging.Limits{
		Logger:        messagingLogger,
		RateLimiter:   limiter,
		UsageLimiter:  usageLimiter,
		RemoteIP:      remoteIP,
		Config:        messagingRateLimitsConfig,
		FeatureConfig: messagingFeatureConfig,
		EnvConfig:     rateLimitsEnvironmentConfig,
	}
	serviceLogger := whatsapp.NewServiceLogger(factory)
	devMode := environmentConfig.DevMode
	testModeWhatsappSuppressed := deps.ProvideTestModeWhatsappSuppressed(testModeFeatureConfig)
	whatsappConfig := messagingConfig.Whatsapp
	whatsappOnPremisesCredentials := deps.ProvideWhatsappOnPremisesCredentials(secretConfig)
	tokenStore := &whatsapp.TokenStore{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	onPremisesClient := whatsapp.NewWhatsappOnPremisesClient(whatsappConfig, whatsappOnPremisesCredentials, tokenStore)
	whatsappService := &whatsapp.Service{
		Context:                    contextContext,
		Logger:                     serviceLogger,
		DevMode:                    devMode,
		TestModeWhatsappSuppressed: testModeWhatsappSuppressed,
		Config:                     whatsappConfig,
		OnPremisesClient:           onPremisesClient,
		TokenStore:                 tokenStore,
	}
	sender := &messaging.Sender{
		Limits:    limits,
		TaskQueue: queue,
		Events:    eventService,
		Whatsapp:  whatsappService,
	}
	messageSender := &otp.MessageSender{
		Translation:     translationService,
		Endpoints:       endpointsEndpoints,
		Sender:          sender,
		WhatsappService: whatsappService,
	}
	oAuthSSOProviderCredentials := deps.ProvideOAuthSSOProviderCredentials(secretConfig)
	normalizer := &stdattrs2.Normalizer{
		LoginIDNormalizerFactory: normalizerFactory,
	}
	oAuthProviderFactory := &sso.OAuthProviderFactory{
		Endpoints:                    endpointsEndpoints,
		IdentityConfig:               identityConfig,
		Credentials:                  oAuthSSOProviderCredentials,
		RedirectURL:                  endpointsEndpoints,
		Clock:                        clockClock,
		WechatURLProvider:            endpointsEndpoints,
		StandardAttributesNormalizer: normalizer,
	}
	forgotpasswordLogger := forgotpassword.NewLogger(factory)
	forgotpasswordService := &forgotpassword.Service{
		Logger:         forgotpasswordLogger,
		Config:         appConfig,
		FeatureConfig:  featureConfig,
		Identities:     identityFacade,
		Authenticators: authenticatorFacade,
		OTPCodes:       otpService,
		OTPSender:      messageSender,
	}
	responseWriter := p.ResponseWriter
	nonceService := &nonce.Service{
		Cookies:        cookieManager,
		Request:        request,
		ResponseWriter: responseWriter,
	}
	challengeProvider := &challenge.Provider{
		Redis: appredisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	authenticationinfoStoreRedis := &authenticationinfo.StoreRedis{
		Context: contextContext,
		Redis:   appredisHandle,
		AppID:   appID,
	}
	manager2 := &session.Manager{
		IDPSessions:         idpsessionManager,
		AccessTokenSessions: sessionManager,
		Events:              eventService,
	}
	mfaCookieDef := mfa.NewDeviceTokenCookieDef(authenticationConfig)
	interactionContext := &interaction.Context{
		Request:                         request,
		RemoteIP:                        remoteIP,
		Database:                        sqlExecutor,
		Clock:                           clockClock,
		Config:                          appConfig,
		FeatureConfig:                   featureConfig,
		OfflineGrants:                   redisStore,
		Identities:                      identityFacade,
		Authenticators:                  authenticatorFacade,
		AnonymousIdentities:             anonymousProvider,
		AnonymousUserPromotionCodeStore: anonymousStoreRedis,
		BiometricIdentities:             biometricProvider,
		OTPCodeService:                  otpService,
		OTPSender:                       messageSender,
		OAuthProviderFactory:            oAuthProviderFactory,
		MFA:                             mfaService,
		ForgotPassword:                  forgotpasswordService,
		ResetPassword:                   forgotpasswordService,
		Passkey:                         passkeyService,
		LoginIDNormalizerFactory:        normalizerFactory,
		Verification:                    verificationService,
		RateLimiter:                     limiter,
		Nonces:                          nonceService,
		Challenges:                      challengeProvider,
		Users:                           userProvider,
		StdAttrsService:                 stdattrsService,
		Events:                          eventService,
		CookieManager:                   cookieManager,
		AuthenticationInfoService:       authenticationinfoStoreRedis,
		Sessions:                        idpsessionProvider,
		SessionManager:                  manager2,
		SessionCookie:                   cookieDef,
		MFADeviceTokenCookie:            mfaCookieDef,
	}
	interactionStoreRedis := &interaction.StoreRedis{
		Redis: appredisHandle,
		AppID: appID,
	}
	interactionService := &interaction.Service{
		Logger:  interactionLogger,
		Context: interactionContext,
		Store:   interactionStoreRedis,
	}
	serviceInteractionService := &service3.InteractionService{
		Graph: interactionService,
	}
	facadeUserFacade := &facade2.UserFacade{
		UserSearchService:  elasticsearchService,
		Users:              userFacade,
		StandardAttributes: serviceNoEvent,
		Interaction:        serviceInteractionService,
	}
	auditLogFeatureConfig := featureConfig.AuditLog
	auditLogFacade := &facade2.AuditLogFacade{
		AuditLogQuery:         query,
		Clock:                 clockClock,
		AuditDatabase:         readHandle,
		AuditLogFeatureConfig: auditLogFeatureConfig,
	}
	facadeIdentityFacade := &facade2.IdentityFacade{
		Identities:  serviceService,
		Interaction: serviceInteractionService,
	}
	facadeAuthenticatorFacade := &facade2.AuthenticatorFacade{
		Authenticators: service4,
		Interaction:    serviceInteractionService,
	}
	adminVerificationFacade := &facade.AdminVerificationFacade{
		Verification: verificationService,
		Coordinator:  coordinator,
	}
	verificationFacade := &facade2.VerificationFacade{
		Verification: adminVerificationFacade,
	}
	sessionFacade := &facade2.SessionFacade{
		Sessions: manager2,
	}
	userProfileFacade := &facade2.UserProfileFacade{
		User:               userFacade,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
		Events:             eventService,
	}
	authorizationService := &oauth2.AuthorizationService{
		AppID:               appID,
		Store:               authorizationStore,
		Clock:               clockClock,
		OAuthSessionManager: sessionManager,
	}
	authorizationFacade := &facade2.AuthorizationFacade{
		Authorizations: authorizationService,
	}
	oAuthKeyMaterials := deps.ProvideOAuthKeyMaterials(secretConfig)
	idTokenIssuer := &oidc.IDTokenIssuer{
		Secrets: oAuthKeyMaterials,
		BaseURL: endpointsEndpoints,
		Users:   queries,
		Clock:   clockClock,
	}
	accessTokenEncoding := &oauth2.AccessTokenEncoding{
		Secrets:    oAuthKeyMaterials,
		Clock:      clockClock,
		UserClaims: idTokenIssuer,
		BaseURL:    endpointsEndpoints,
		Events:     eventService,
	}
	tokenGenerator := _wireTokenGeneratorValue
	tokenService := &handler.TokenService{
		RemoteIP:            remoteIP,
		UserAgentString:     userAgentString,
		AppID:               appID,
		Config:              oAuthConfig,
		Authorizations:      authorizationStore,
		OfflineGrants:       redisStore,
		AccessGrants:        redisStore,
		OfflineGrantService: offlineGrantService,
		AccessEvents:        eventProvider,
		AccessTokenIssuer:   accessTokenEncoding,
		GenerateToken:       tokenGenerator,
		Clock:               clockClock,
		Users:               queries,
	}
	oAuthFacade := &facade2.OAuthFacade{
		Config:         oAuthConfig,
		Users:          userFacade,
		Authorizations: authorizationService,
		Tokens:         tokenService,
		Clock:          clockClock,
	}
	oauthOfflineGrantService := &oauth2.OfflineGrantService{
		OAuthConfig: oAuthConfig,
		Clock:       clockClock,
		IDPSessions: idpsessionProvider,
	}
	sessionListingService := &sessionlisting.SessionListingService{
		OAuthConfig:   oAuthConfig,
		IDPSessions:   idpsessionProvider,
		OfflineGrants: oauthOfflineGrantService,
	}
	graphqlContext := &graphql.Context{
		GQLLogger:             logger,
		Config:                appConfig,
		OAuthConfig:           oAuthConfig,
		AdminAPIFeatureConfig: adminAPIFeatureConfig,
		Users:                 userLoader,
		Identities:            identityLoader,
		Authenticators:        authenticatorLoader,
		AuditLogs:             auditLogLoader,
		UserFacade:            facadeUserFacade,
		AuditLogFacade:        auditLogFacade,
		IdentityFacade:        facadeIdentityFacade,
		AuthenticatorFacade:   facadeAuthenticatorFacade,
		VerificationFacade:    verificationFacade,
		SessionFacade:         sessionFacade,
		UserProfileFacade:     userProfileFacade,
		AuthorizationFacade:   authorizationFacade,
		OAuthFacade:           oAuthFacade,
		SessionListing:        sessionListingService,
		OTPCode:               otpService,
		ForgotPassword:        forgotpasswordService,
		Events:                eventService,
	}
	graphQLHandler := &transport.GraphQLHandler{
		GraphQLContext: graphqlContext,
		AppDatabase:    handle,
	}
	return graphQLHandler
}

var (
	_wireRandValue           = idpsession.Rand(rand.SecureRand)
	_wireTokenGeneratorValue = handler.TokenGenerator(oauth2.GenerateToken)
)

func newPresignImagesUploadHandler(p *deps.RequestProvider) http.Handler {
	appProvider := p.AppProvider
	factory := appProvider.LoggerFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	request := p.Request
	rootProvider := appProvider.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	appContext := appProvider.AppContext
	configConfig := appContext.Config
	appConfig := configConfig.AppConfig
	appID := appConfig.ID
	secretConfig := configConfig.SecretConfig
	imagesKeyMaterials := deps.ProvideImagesKeyMaterials(secretConfig)
	clockClock := _wireSystemClockValue
	provider := &presign.Provider{
		Secret: imagesKeyMaterials,
		Clock:  clockClock,
		Host:   httpHost,
	}
	presignImagesUploadHandlerLogger := transport.NewPresignImagesUploadHandlerLogger(factory)
	presignImagesUploadHandler := &transport.PresignImagesUploadHandler{
		JSON:            jsonResponseWriter,
		HTTPProto:       httpProto,
		HTTPHost:        httpHost,
		AppID:           appID,
		PresignProvider: provider,
		Logger:          presignImagesUploadHandlerLogger,
	}
	return presignImagesUploadHandler
}
