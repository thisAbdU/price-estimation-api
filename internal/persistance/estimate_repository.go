package persistance

import (
	"context"
	"database/sql"
	"price-estimation-api/internal/db"
	"price-estimation-api/internal/domain"
	"price-estimation-api/internal/ports"
)

type estimateRepository struct {
    conn    *sql.DB
    queries *db.Queries
}

func NewEstimateRepository(conn *sql.DB) ports.EstimateRepository {
    return &estimateRepository{
        conn:    conn,
        queries: db.New(conn),
    }
}

func (r *estimateRepository) CreateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error){
    createdEstimate, err := r.queries.CreateEstimate(ctx, db.CreateEstimateParams{
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    })
    
    if err != nil {
        return domain.Estimate{}, err
    }

    return createdEstimate, nil
}

func (r *estimateRepository) GetEstimates(ctx context.Context) ([]domain.Estimate, error) {
    estimates, err := r.queries.GetEstimates(ctx)
    if err != nil {
        return nil, err
    }

    return estimates, nil
}

func (r *estimateRepository) GetEstimateByID(ctx context.Context, id int) (domain.Estimate, error) {
    estimate, err := r.queries.GetEstimateByID(ctx, db.GetEstimateByIDParams{ID: id})

    if err != nil {
        return domain.Estimate{}, err
    }

    return estimate, nil
}

func (r *estimateRepository) UpdateEstimate(ctx context.Context, estimate domain.Estimate) (domain.Estimate, error) {
    updatedEstimate, err := r.queries.UpdateEstimate(ctx, db.UpdateEstimateParams{
        ID:          estimate.ID,
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    })
    if err != nil {
        return domain.Estimate{}, err
    }

    return  updatedEstimate, err
}

func (r *estimateRepository) DeleteEstimate(ctx context.Context, id int) error {
    err := r.queries.DeleteEstimate(ctx, db.DeleteEstimateParams{ID: id})
    if err != nil {
        return err
    }

    return nil

}
