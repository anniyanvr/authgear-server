// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmdpricing

import (
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	config2 "github.com/authgear/authgear-server/pkg/portal/config"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/cobrasentry"
	"github.com/getsentry/sentry-go"
)

// Injectors from wire.go:

func NewStripeService(pool *db.Pool, databaseCredentials *config.DatabaseCredentials, stripeConfig *config2.StripeConfig, hub *sentry.Hub) *StripeService {
	factory := cobrasentry.NewLoggerFactory(hub)
	logger := NewLogger(factory)
	api := NewClientAPI(stripeConfig, logger)
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	handle := globaldb.NewHandle(pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(handle)
	store := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	clock := _wireSystemClockValue
	stripeService := &StripeService{
		ClientAPI:   api,
		Handle:      handle,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Store:       store,
		Clock:       clock,
		Logger:      logger,
	}
	return stripeService
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)
