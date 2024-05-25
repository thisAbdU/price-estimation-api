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


type Service struct {
	estimateRepo  ports.EstimateRepository
}

func NewService(repo ports.EstimateRepository) *Service {
	return &Service{estimateRepo: repo}
}

func (s *Service) CreateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error) {
	return s.estimateRepo.CreateEstimate(ctx, estimate)
}

func (s *Service) GetEstimates(ctx context.Context) ([]domain.Estimate, error) {
	return s.estimateRepo.GetEstimates(ctx)
}

func (s *Service) GetEstimate(ctx context.Context, id int) (domain.Estimate, error) {
	return s.estimateRepo.GetEstimateByID(ctx, id)
}

func (s *Service) DeleteEstimate(ctx context.Context, id int) error {
	return s.estimateRepo.DeleteEstimate(ctx, id)
}

func (s *Service) UpdateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error) {
	return s.estimateRepo.UpdateEstimate(ctx, estimate)
}