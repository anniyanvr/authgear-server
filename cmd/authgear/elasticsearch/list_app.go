package elasticsearch

import (
	"context"

	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
)

type AppLister struct {
	Handle *globaldb.Handle
	Store  *configsource.Store
}

func (l *AppLister) ListApps(ctx context.Context) (appIDs []string, err error) {
	err = l.Handle.ReadOnly(ctx, func(ctx context.Context) error {
		srcs, err := l.Store.ListAll(ctx)
		if err != nil {
			return err
		}
		for _, src := range srcs {
			appID := src.AppID
			appIDs = append(appIDs, appID)
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
