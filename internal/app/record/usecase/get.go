package usecase

import (
	"context"

	ls_d "github.com/bartmika/logfarm/internal/domain/record"
)

func (uc recordUsecase) GetByID(ctx context.Context, id string) (*ls_d.Record, error) {
	return uc.RecordRepo.GetByID(ctx, id)
}
