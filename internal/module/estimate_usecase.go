package module

import (
	"context"
	"price-estimation-api/internal/db"
	"price-estimation-api/internal/ports"
)

type EstimateUsecase interface {
	CreateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error)
	GetEstimates(ctx context.Context, limit, offset int32) ([]db.Estimate, error)
	GetEstimate(ctx context.Context, id int32) (db.Estimate, error)
	DeleteEstimate(ctx context.Context, id int32) error
	UpdateEstimate(ctx context.Context, id int32) (db.Estimate, error)
}

type estimateUsecase struct {
	estimateRepo ports.EstimateRepository
}

// UpdateEstimate implements EstimateUsecase.
func (e *estimateUsecase) UpdateEstimate(ctx context.Context, id int32) (db.Estimate, error) {
	panic("unimplemented")
}

func NewEstimateUsecase(estimateRepo ports.EstimateRepository) EstimateUsecase {
	return &estimateUsecase{
		estimateRepo: estimateRepo,
	}
}

func (e *estimateUsecase) CreateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error) {
	return e.estimateRepo.CreateEstimate(ctx, estimate)
}

func (e *estimateUsecase) GetEstimates(ctx context.Context, limit, offset int32) ([]db.Estimate, error) {
	return e.estimateRepo.GetEstimatesWithPagination(ctx, limit, offset)
}

func (e *estimateUsecase) GetEstimate(ctx context.Context, id int32) (db.Estimate, error) {
	return e.estimateRepo.GetEstimateByID(ctx, id)
}

func (e *estimateUsecase) DeleteEstimate(ctx context.Context, id int32) error {
	return e.estimateRepo.DeleteEstimate(ctx, id)
}
