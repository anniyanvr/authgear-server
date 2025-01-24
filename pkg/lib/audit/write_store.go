package audit

import (
	"context"
	"encoding/json"

	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
)

type WriteStore struct {
	SQLBuilder  *auditdb.SQLBuilderApp
	SQLExecutor *auditdb.WriteSQLExecutor
}

func (s *WriteStore) PersistLog(ctx context.Context, logEntry *Log) (err error) {
	data, err := json.Marshal(logEntry.Data)
	if err != nil {
		return
	}

	builder := s.SQLBuilder.
		Insert(s.SQLBuilder.TableName("_audit_log")).
		Columns(
			"id",
			"created_at",
			"user_id",
			"activity_type",
			"ip_address",
			"user_agent",
			"client_id",
			"data",
		).
		Values(
			logEntry.ID,
			logEntry.CreatedAt,
			logEntry.UserID,
			logEntry.ActivityType,
			logEntry.IPAddress,
			logEntry.UserAgent,
			logEntry.ClientID,
			data,
		)

	_, err = s.SQLExecutor.ExecWith(ctx, builder)
	if err != nil {
		return
	}

	return nil
}
