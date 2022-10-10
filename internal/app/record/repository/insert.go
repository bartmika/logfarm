package repository

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/logfarm/internal/domain/record"
)

func (r *RecordRepoImpl) Insert(ctx context.Context, m *domain.Record) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sql, args, err := sq.
		Insert("records").
		Columns(
			"id",
			"client",
			"content",
			"facility",
			"hostname",
			"priority",
			"severity",
			"tag",
			"timestamp",
			"tls_peer",
		).
		Values(
			m.ID,
			m.Client,
			m.Content,
			m.Facility,
			m.Hostname,
			m.Priority,
			m.Severity,
			m.Tag,
			m.Timestamp,
			m.TLSPeer,
		).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.Println("RecordRepoImpl | Insert | err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	return err
}
