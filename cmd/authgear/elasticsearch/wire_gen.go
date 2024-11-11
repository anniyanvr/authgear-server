// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package elasticsearch

import (
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/rolesgroups"
	"github.com/authgear/authgear-server/pkg/util/clock"
)

// Injectors from wire.go:

func NewAppLister(pool *db.Pool, databaseCredentials *config.DatabaseCredentials) *AppLister {
	globalDatabaseCredentialsEnvironmentConfig := NewGlobalDatabaseCredentials(databaseCredentials)
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := globaldb.NewHandle(pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(handle)
	store := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	appLister := &AppLister{
		Handle: handle,
		Store:  store,
	}
	return appLister
}

func NewReindexer(pool *db.Pool, databaseCredentials *config.DatabaseCredentials, appID config.AppID) *Reindexer {
	clock := _wireSystemClockValue
	databaseEnvironmentConfig := config.NewDefaultDatabaseEnvironmentConfig()
	factory := NewLoggerFactory()
	handle := appdb.NewHandle(pool, databaseEnvironmentConfig, databaseCredentials, factory)
	sqlBuilderApp := appdb.NewSQLBuilderApp(databaseCredentials, appID)
	sqlExecutor := appdb.NewSQLExecutor(handle)
	store := &user.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clock,
		AppID:       appID,
	}
	identityConfig := NewEmptyIdentityConfig()
	oauthStore := &oauth.Store{
		SQLBuilder:     sqlBuilderApp,
		SQLExecutor:    sqlExecutor,
		IdentityConfig: identityConfig,
	}
	loginidStore := &loginid.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	rolesgroupsStore := &rolesgroups.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clock,
	}
	reindexedTimestamps := NewReindexedTimestamps()
	reindexer := &Reindexer{
		Clock:               clock,
		Handle:              handle,
		AppID:               appID,
		Users:               store,
		OAuth:               oauthStore,
		LoginID:             loginidStore,
		RolesGroups:         rolesgroupsStore,
		ReindexedTimestamps: reindexedTimestamps,
	}
	return reindexer
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)
