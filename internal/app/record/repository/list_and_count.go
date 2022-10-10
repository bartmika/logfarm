package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/logfarm/internal/domain/record"
)

func (s *RecordRepoImpl) getWhereKeysByFilter(f *domain.RecordFilter) sq.And {
	// Apply specific 'where' keys to apply.
	k := sq.And{}

	if f.Tag != "" {
		k = append(k, sq.Eq{"tag": f.Tag})
	}

	if len(f.IDs) > 0 {
		idKeys := sq.Or{}
		for _, id := range f.IDs { // https://github.com/Masterminds/squirrel/issues/104#issuecomment-475309510
			idKeys = append(idKeys, sq.Eq{"id": id})
		}
		k = append(k, idKeys)
		// log.Println("k:", k)
	}

	if !f.TimestampGreaterThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Gt{"timestamp": f.TimestampGreaterThenOrEqual},
			sq.Eq{"timestamp": f.TimestampGreaterThenOrEqual},
		})
	}
	if !f.TimestampGreaterThen.IsZero() {
		k = append(k, sq.Gt{"timestamp": f.TimestampGreaterThen})
	}
	if !f.TimestampLessThen.IsZero() {
		k = append(k, sq.Lt{"timestamp": f.TimestampLessThen})
	}
	if !f.TimestampLessThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Lt{"timestamp": f.TimestampLessThenOrEqual},
			sq.Eq{"timestamp": f.TimestampLessThenOrEqual},
		})
	}

	return k
}

func (s *RecordRepoImpl) ListByFilter(ctx context.Context, f *domain.RecordFilter) ([]*domain.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
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
	).From("records")
	k := s.getWhereKeysByFilter(f)

	rds = rds.Where(k)

	rds = rds.OrderBy(f.SortField + " " + f.SortOrder)

	// Note:
	// (1) https://ivopereira.net/efficient-pagination-dont-use-offset-limit
	// (2) https://github.com/Masterminds/squirrel/blob/def598cbb358368fbfc3f6a9a914699a36846992/select_test.go#L41

	// rds = rds.Offset(f.Offset).Suffix("FETCH FIRST ? ROWS ONLY", f.Limit)

	// Build the SQL statement and the accomponing arguments.
	sql, args, err := rds.ToSql()

	// // For debugging purposes only.
	// log.Println("f:", f)
	// log.Println("sql:", sql)
	// log.Println("args:", args)
	// log.Println("err:", err)

	stmt, err := s.db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	var arr []*domain.Record
	defer rows.Close()
	for rows.Next() {
		m := new(domain.Record)
		err := rows.Scan(
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
			return nil, err
		}
		arr = append(arr, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if arr == nil {
		return []*domain.Record{}, nil
	}
	return arr, err
}

func (s *RecordRepoImpl) CountByFilter(ctx context.Context, f *domain.RecordFilter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// The result we are looking for.
	var count uint64

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	submissionCount := psql.Select(
		"count(*)",
	).From("records")

	k := s.getWhereKeysByFilter(f)

	submissionCount = submissionCount.Where(k)

	// Build the SQL statement and the accomponing arguments.
	sql, args, err := submissionCount.ToSql()

	err = s.db.QueryRowContext(ctx, sql, args...).Scan(&count)
	return count, err
}
