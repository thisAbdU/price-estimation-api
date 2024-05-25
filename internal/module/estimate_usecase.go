package module

import (
	"context"
	"price-estimation-api/internal/domain"
	"price-estimation-api/internal/ports"

)

type EstimateUsecase interface {
	CreateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error)
	GetEstimates(ctx context.Context) ([]domain.Estimate, error)
	GetEstimate(ctx context.Context, id int) (domain.Estimate, error)
	DeleteEstimate(ctx context.Context, id int) error
	UpdateEstimate(ctx context.Context, id int) (domain.Estimate, error)
}


type estimateUsecase struct {
	estimateRepo  ports.EstimateRepository
}

func NewEstimateUsecase(repo ports.EstimateRepository) EstimateUsecase {
	return &estimateUsecase{estimateRepo: repo}
}

func (s *estimateUsecase) CreateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error) {
	return s.estimateRepo.CreateEstimate(ctx, estimate)
}

func (s *estimateUsecase) GetEstimates(ctx context.Context) ([]domain.Estimate, error) {
	return s.estimateRepo.GetEstimates(ctx)
}

func (s *estimateUsecase) GetEstimate(ctx context.Context, id int) (domain.Estimate, error) {
	return s.estimateRepo.GetEstimateByID(ctx, id)
}

func (s *estimateUsecase) DeleteEstimate(ctx context.Context, id int) error {
	return s.estimateRepo.DeleteEstimate(ctx, id)
}

func (s *estimateUsecase) UpdateEstimate(ctx context.Context, id int) (domain.Estimate, error) {
	return s.estimateRepo.UpdateEstimate(ctx, id)
	
}