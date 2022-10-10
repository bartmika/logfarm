package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/logfarm/internal/domain/record"
)

func (r *RecordRepoImpl) getBy(ctx context.Context, k *sq.And) (*domain.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select(
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
		From("records").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m := new(domain.Record)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&m.ID,
		&m.Client,
		&m.Content,
		&m.Facility,
		&m.Hostname,
		&m.Priority,
		&m.Severity,
		&m.Tag,
		&m.Timestamp,
		&m.TLSPeer,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// CASE 2 OF 2: All other errors.
		return nil, err
	}

	return m, nil
}

func (r *RecordRepoImpl) GetByID(ctx context.Context, id string) (*domain.Record, error) {
	k := &sq.And{
		sq.Eq{"id": id},
	}
	return r.getBy(ctx, k)
}

func (r *RecordRepoImpl) GetByUUID(ctx context.Context, uid string) (*domain.Record, error) {
	k := &sq.And{
		sq.Eq{"uuid": uid},
	}
	return r.getBy(ctx, k)
}
