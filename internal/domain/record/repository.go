package record

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, u *Record) error
	GetByID(ctx context.Context, id string) (*Record, error)
	ListByFilter(ctx context.Context, filter *RecordFilter) ([]*Record, error)
	CountByFilter(ctx context.Context, filter *RecordFilter) (uint64, error)
	UpdateByID(ctx context.Context, u *Record) error
	CheckIfExistsByID(ctx context.Context, id string) (bool, error)
	InsertOrUpdateByID(ctx context.Context, u *Record) error
	DeleteByID(ctx context.Context, id string) error
}
