// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package background

import (
	"context"
	"github.com/authgear/authgear-server/pkg/lib/audit"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/oob"
	passkey3 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/passkey"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/password"
	service2 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/totp"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/anonymous"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/biometric"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/passkey"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/siwe"
	"github.com/authgear/authgear-server/pkg/lib/authn/mfa"
	"github.com/authgear/authgear-server/pkg/lib/authn/otp"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/elasticsearch"
	"github.com/authgear/authgear-server/pkg/lib/event"
	"github.com/authgear/authgear-server/pkg/lib/facade"
	"github.com/authgear/authgear-server/pkg/lib/feature/accountanonymization"
	"github.com/authgear/authgear-server/pkg/lib/feature/accountdeletion"
	"github.com/authgear/authgear-server/pkg/lib/feature/customattrs"
	passkey2 "github.com/authgear/authgear-server/pkg/lib/feature/passkey"
	siwe2 "github.com/authgear/authgear-server/pkg/lib/feature/siwe"
	"github.com/authgear/authgear-server/pkg/lib/feature/stdattrs"
	"github.com/authgear/authgear-server/pkg/lib/feature/verification"
	"github.com/authgear/authgear-server/pkg/lib/feature/web3"
	"github.com/authgear/authgear-server/pkg/lib/feature/welcomemessage"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	oauth2 "github.com/authgear/authgear-server/pkg/lib/oauth"
	"github.com/authgear/authgear-server/pkg/lib/oauth/pq"
	"github.com/authgear/authgear-server/pkg/lib/oauth/redis"
	"github.com/authgear/authgear-server/pkg/lib/ratelimit"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/lib/session/access"
	"github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/lib/tutorial"
	"github.com/authgear/authgear-server/pkg/lib/web"
	"github.com/authgear/authgear-server/pkg/util/backgroundjob"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/rand"
	"github.com/authgear/authgear-server/pkg/util/template"
)

// Injectors from wire.go:

func newConfigSourceController(p *deps.BackgroundProvider, c context.Context) *configsource.Controller {
	config := p.ConfigSourceConfig
	factory := p.LoggerFactory
	localFSLogger := configsource.NewLocalFSLogger(factory)
	manager := p.BaseResources
	localFS := &configsource.LocalFS{
		Logger:        localFSLogger,
		BaseResources: manager,
		Config:        config,
	}
	databaseLogger := configsource.NewDatabaseLogger(factory)
	environmentConfig := p.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	clock := _wireSystemClockValue
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	pool := p.DatabasePool
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(c, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlExecutor := globaldb.NewSQLExecutor(c, handle)
	store := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	resolveAppIDType := configsource.NewResolveAppIDTypeDomain()
	database := &configsource.Database{
		Logger:              databaseLogger,
		BaseResources:       manager,
		TrustProxy:          trustProxy,
		Config:              config,
		Clock:               clock,
		Store:               store,
		Database:            handle,
		DatabaseCredentials: globalDatabaseCredentialsEnvironmentConfig,
		DatabaseConfig:      databaseEnvironmentConfig,
		ResolveAppIDType:    resolveAppIDType,
	}
	controller := configsource.NewController(config, localFS, database)
	return controller
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)

func newAccountAnonymizationRunner(p *deps.BackgroundProvider, c context.Context, ctrl *configsource.Controller) *backgroundjob.Runner {
	factory := p.LoggerFactory
	pool := p.DatabasePool
	environmentConfig := p.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(c, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(c, handle)
	clockClock := _wireSystemClockValue
	store := &accountanonymization.Store{
		Handle:      handle,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	accountAnonymizationServiceFactory := &AccountAnonymizationServiceFactory{
		BackgroundProvider: p,
	}
	runnableLogger := accountanonymization.NewRunnableLogger(factory)
	runnable := &accountanonymization.Runnable{
		Store:              store,
		AppContextResolver: ctrl,
		UserServiceFactory: accountAnonymizationServiceFactory,
		Logger:             runnableLogger,
	}
	runner := accountanonymization.NewRunner(factory, runnable)
	return runner
}

func newAccountDeletionRunner(p *deps.BackgroundProvider, c context.Context, ctrl *configsource.Controller) *backgroundjob.Runner {
	factory := p.LoggerFactory
	pool := p.DatabasePool
	environmentConfig := p.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(c, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(c, handle)
	clockClock := _wireSystemClockValue
	store := &accountdeletion.Store{
		Handle:      handle,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	accountDeletionServiceFactory := &AccountDeletionServiceFactory{
		BackgroundProvider: p,
	}
	runnableLogger := accountdeletion.NewRunnableLogger(factory)
	runnable := &accountdeletion.Runnable{
		Store:              store,
		AppContextResolver: ctrl,
		UserServiceFactory: accountDeletionServiceFactory,
		Logger:             runnableLogger,
	}
	runner := accountdeletion.NewRunner(factory, runnable)
	return runner
}

func newUserService(ctx context.Context, p *deps.BackgroundProvider, appID string, appContext *config.AppContext) *UserService {
	pool := p.DatabasePool
	environmentConfig := p.EnvironmentConfig
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	configConfig := appContext.Config
	secretConfig := configConfig.SecretConfig
	databaseCredentials := deps.ProvideDatabaseCredentials(secretConfig)
	factory := p.LoggerFactory
	handle := appdb.NewHandle(ctx, pool, databaseEnvironmentConfig, databaseCredentials, factory)
	appConfig := configConfig.AppConfig
	configAppID := appConfig.ID
	sqlBuilderApp := appdb.NewSQLBuilderApp(databaseCredentials, configAppID)
	sqlExecutor := appdb.NewSQLExecutor(ctx, handle)
	clockClock := _wireSystemClockValue
	store := &user.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	manager := p.BaseResources
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
	httpProto := ProvideHTTPProto()
	webAppCDNHost := environmentConfig.WebAppCDNHost
	globalEmbeddedResourceManager := p.EmbeddedResources
	staticAssetResolver := &web.StaticAssetResolver{
		Context:           ctx,
		Config:            httpConfig,
		Localization:      localizationConfig,
		HTTPProto:         httpProto,
		WebAppCDNHost:     webAppCDNHost,
		Resources:         manager,
		EmbeddedResources: globalEmbeddedResourceManager,
	}
	translationService := &translation.Service{
		Context:        ctx,
		TemplateEngine: engine,
		StaticAssets:   staticAssetResolver,
	}
	logger := ratelimit.NewLogger(factory)
	redisPool := p.RedisPool
	hub := p.RedisHub
	redisEnvironmentConfig := &environmentConfig.RedisConfig
	redisCredentials := deps.ProvideRedisCredentials(secretConfig)
	appredisHandle := appredis.NewHandle(redisPool, hub, redisEnvironmentConfig, redisCredentials, factory)
	storageRedis := &ratelimit.StorageRedis{
		AppID: configAppID,
		Redis: appredisHandle,
	}
	featureConfig := configConfig.FeatureConfig
	rateLimitsFeatureConfig := featureConfig.RateLimits
	limiter := &ratelimit.Limiter{
		Logger:  logger,
		Storage: storageRedis,
		Clock:   clockClock,
		Config:  rateLimitsFeatureConfig,
	}
	welcomeMessageConfig := appConfig.WelcomeMessage
	noopTaskQueue := NewNoopTaskQueue()
	remoteIP := ProvideRemoteIP()
	userAgentString := ProvideUserAgentString()
	eventLogger := event.NewLogger(factory)
	sqlBuilder := appdb.NewSQLBuilder(databaseCredentials)
	storeImpl := &event.StoreImpl{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
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
	store2 := &passkey2.Store{
		Context: ctx,
		Redis:   appredisHandle,
		AppID:   configAppID,
	}
	request := NewDummyHTTPRequest()
	trustProxy := environmentConfig.TrustProxy
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
	web3Config := appConfig.Web3
	storeRedis := &siwe2.StoreRedis{
		Context: ctx,
		Redis:   appredisHandle,
		AppID:   configAppID,
		Clock:   clockClock,
	}
	siweLogger := siwe2.NewLogger(factory)
	siweService := &siwe2.Service{
		RemoteIP:    remoteIP,
		HTTPConfig:  httpConfig,
		Web3Config:  web3Config,
		Clock:       clockClock,
		NonceStore:  storeRedis,
		RateLimiter: limiter,
		Logger:      siweLogger,
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
	codeStoreRedis := &otp.CodeStoreRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	lookupStoreRedis := &otp.LookupStoreRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	attemptTrackerRedis := &otp.AttemptTrackerRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	otpLogger := otp.NewLogger(factory)
	otpService := &otp.Service{
		Clock:          clockClock,
		AppID:          configAppID,
		RemoteIP:       remoteIP,
		CodeStore:      codeStoreRedis,
		LookupStore:    lookupStoreRedis,
		AttemptTracker: attemptTrackerRedis,
		Logger:         otpLogger,
		RateLimiter:    limiter,
	}
	rateLimits := service2.RateLimits{
		IP:          remoteIP,
		Config:      authenticationConfig,
		RateLimiter: limiter,
	}
	service3 := &service2.Service{
		Store:          store3,
		Config:         appConfig,
		Password:       passwordProvider,
		Passkey:        provider2,
		TOTP:           totpProvider,
		OOBOTP:         oobProvider,
		OTPCodeService: otpService,
		RateLimits:     rateLimits,
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
	httpHost := ProvideHTTPHost()
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
		Authenticators:     service3,
		Verification:       verificationService,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
		Web3:               web3Service,
	}
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
		Context:         ctx,
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
	auditDatabaseCredentials := deps.ProvideAuditDatabaseCredentials(secretConfig)
	writeHandle := auditdb.NewWriteHandle(ctx, pool, databaseEnvironmentConfig, auditDatabaseCredentials, factory)
	auditdbSQLBuilderApp := auditdb.NewSQLBuilderApp(auditDatabaseCredentials, configAppID)
	writeSQLExecutor := auditdb.NewWriteSQLExecutor(ctx, writeHandle)
	writeStore := &audit.WriteStore{
		SQLBuilder:  auditdbSQLBuilderApp,
		SQLExecutor: writeSQLExecutor,
	}
	auditSink := &audit.Sink{
		Logger:   auditLogger,
		Database: writeHandle,
		Store:    writeStore,
	}
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	globaldbSQLBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	globaldbHandle := globaldb.NewHandle(ctx, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	globaldbSQLExecutor := globaldb.NewSQLExecutor(ctx, globaldbHandle)
	tutorialStoreImpl := &tutorial.StoreImpl{
		SQLBuilder:  globaldbSQLBuilder,
		SQLExecutor: globaldbSQLExecutor,
	}
	tutorialService := &tutorial.Service{
		Store: tutorialStoreImpl,
	}
	tutorialSink := &tutorial.Sink{
		AppID:        configAppID,
		Service:      tutorialService,
		GlobalHandle: globaldbHandle,
	}
	elasticsearchLogger := elasticsearch.NewLogger(factory)
	elasticsearchCredentials := deps.ProvideElasticsearchCredentials(secretConfig)
	client := elasticsearch.NewClient(elasticsearchCredentials)
	elasticsearchService := elasticsearch.Service{
		AppID:     configAppID,
		Client:    client,
		Users:     queries,
		OAuth:     oauthStore,
		LoginID:   loginidStore,
		TaskQueue: noopTaskQueue,
	}
	elasticsearchSink := &elasticsearch.Sink{
		Logger:   elasticsearchLogger,
		Service:  elasticsearchService,
		Database: handle,
	}
	eventService := event.NewService(ctx, remoteIP, userAgentString, eventLogger, handle, clockClock, localizationConfig, storeImpl, resolverImpl, sink, auditSink, tutorialSink, elasticsearchSink)
	welcomemessageProvider := &welcomemessage.Provider{
		Translation:          translationService,
		RateLimiter:          limiter,
		WelcomeMessageConfig: welcomeMessageConfig,
		TaskQueue:            noopTaskQueue,
		Events:               eventService,
	}
	rawCommands := &user.RawCommands{
		Store:                  store,
		Clock:                  clockClock,
		WelcomeMessageProvider: welcomemessageProvider,
	}
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
		AppID: configAppID,
		Clock: clockClock,
	}
	storeRecoveryCodePQ := &mfa.StoreRecoveryCodePQ{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	mfaService := &mfa.Service{
		IP:            remoteIP,
		DeviceTokens:  storeDeviceTokenRedis,
		RecoveryCodes: storeRecoveryCodePQ,
		Clock:         clockClock,
		Config:        authenticationConfig,
		RateLimiter:   limiter,
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
		AppID:  configAppID,
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
		Context:     ctx,
		Redis:       appredisHandle,
		AppID:       configAppID,
		Logger:      redisLogger,
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	oAuthConfig := appConfig.OAuth
	eventStoreRedis := &access.EventStoreRedis{
		Redis: appredisHandle,
		AppID: configAppID,
	}
	eventProvider := &access.EventProvider{
		Store: eventStoreRedis,
	}
	rand := _wireRandValue
	idpsessionProvider := &idpsession.Provider{
		Context:         ctx,
		RemoteIP:        remoteIP,
		UserAgentString: userAgentString,
		AppID:           configAppID,
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
		Authenticators:             service3,
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
	userService := &UserService{
		AppDBHandle: handle,
		UserFacade:  userFacade,
	}
	return userService
}

var (
	_wireRandValue = idpsession.Rand(rand.SecureRand)
)
