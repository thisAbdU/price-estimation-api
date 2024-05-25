package ports

import (
	"context"
	"price-estimation-api/internal/domain"
)

type EstimateRepository interface {
    CreateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error)
    GetEstimates(ctx context.Context) ([]domain.Estimate, error)
    GetEstimateByID(ctx context.Context, id int) (domain.Estimate, error)
    UpdateEstimate(ctx context.Context, id int) (domain.Estimate, error)
    DeleteEstimate(ctx context.Context, id int) error
}
