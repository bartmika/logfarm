package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/logfarm/internal/domain/record"
)

func (r *RecordRepoImpl) UpdateByID(ctx context.Context, m *domain.Record) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	sql, args, err := sq.
		Update("records").
		Set("client", m.Client).
		Set("content", m.Content).
		Set("facility", m.Facility).
		Set("hostname", m.Hostname).
		Set("priority", m.Priority).
		Set("severity", m.Severity).
		Set("tag", m.Tag).
		Set("timestamp", m.Timestamp).
		Set("tls_peer", m.TLSPeer).
		Where(sq.Eq{"id": m.ID}).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	return err
}
