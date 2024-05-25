package persistance

import (
	"context"
	"database/sql"
	"price-estimation-api/internal/db"
)

type estimateRepository struct {
    queries *db.Queries
}

// CreateEstimate implements ports.EstimateRepository.
func (e *estimateRepository) CreateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error) {
    // Map the input domain.Estimate to db.CreateEstimateParams
    params := db.CreateEstimateParams{
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    }

    // Call the generated CreateEstimate method
    dbEstimate, err := e.queries.CreateEstimate(ctx, params)
    if err != nil {
        return db.Estimate{}, err
    }

    // Return the result
    return dbEstimate, nil
}

// DeleteEstimate implements ports.EstimateRepository.
func (e *estimateRepository) DeleteEstimate(ctx context.Context, id int32) error {
    return e.queries.DeleteEstimate(ctx, id)
}

// GetEstimateByID implements ports.EstimateRepository.
func (e *estimateRepository) GetEstimateByID(ctx context.Context, id int32) (db.Estimate, error) {
    estimate, err := e.queries.GetEstimateByID(ctx, id)
    if err != nil {
        return db.Estimate{}, err
    }
    return estimate, nil
}

// GetEstimatesWithPagination implements ports.EstimateRepository.
func (e *estimateRepository) GetEstimatesWithPagination(ctx context.Context, limit, offset int32) ([]db.Estimate, error) {
    params := db.GetEstimatesWithPaginationParams{
        Limit:  limit,
        Offset: offset,
    }

    estimates, err := e.queries.GetEstimatesWithPagination(ctx, params)
    if err != nil {
        return nil, err
    }

    return estimates, nil
}

// UpdateEstimate implements ports.EstimateRepository.
func (e *estimateRepository) UpdateEstimate(ctx context.Context, estimate db.Estimate) (db.Estimate, error) {
    params := db.UpdateEstimateParams{
        ID:          estimate.ID,
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    }

    updatedEstimate, err := e.queries.UpdateEstimate(ctx, params)
    if err != nil {
        return db.Estimate{}, err
    }

    return updatedEstimate, nil
}

func NewEstimateRepository(dbConn *sql.DB) *estimateRepository {
    return &estimateRepository{
        queries: db.New(dbConn), 
    }
}