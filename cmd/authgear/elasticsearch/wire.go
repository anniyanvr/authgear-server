//go:build wireinject
// +build wireinject

package elasticsearch

import (
	"github.com/google/wire"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
)

func NewAppLister(
	pool *db.Pool,
	databaseCredentials *config.DatabaseCredentials,
) *AppLister {
	panic(wire.Build(DependencySet))
}

func NewReindexer(
	pool *db.Pool,
	databaseCredentials *config.DatabaseCredentials,
	appID config.AppID,
) *Reindexer {
	panic(wire.Build(DependencySet))
}
