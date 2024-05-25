package ports

import (
	"context"
	"price-estimation-api/internal/db"
)

type EstimateRepository interface {
    CreateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error)
    GetEstimatesWithPagination(ctx context.Context, limit, offset int32) ([]db.Estimate, error)
    GetEstimateByID(ctx context.Context, id int32) (db.Estimate, error)
    UpdateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error)
    DeleteEstimate(ctx context.Context, id int32) error
}
