package db

import (
    "context"
    "database/sql"
    "price-estimation-api/internal/domain"
	"price-estimation-api/internal/db"
    "price-estimation-api/internal/ports"
)

type estimateRepository struct {
    conn *sql.DB
    queries *db.Queries
}

func NewEstimateRepository(conn *sql.DB) ports.EstimateRepository {
    return &estimateRepository{
        conn: conn,
        queries: db.New(conn),
    }
}

func (r *estimateRepository) CreateEstimate(estimate domain.Estimate) (domain.Estimate, error) {
    newEstimate, err := r.queries.CreateEstimate(context.Background(), db.CreateEstimateParams{
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    })
    if err != nil {
        return domain.Estimate{}, err
    }

    return domain.Estimate{
        ID:          newEstimate.ID,
        ProductName: newEstimate.ProductName,
        Price:       newEstimate.Price,
        Longitude:   newEstimate.Longitude,
        Latitude:    newEstimate.Latitude,
    }, nil
}

func (r *estimateRepository) GetEstimates() ([]domain.Estimate, error) {
    estimates, err := r.queries.GetEstimates(context.Background())
    if err != nil {
        return nil, err
    }

    var result []domain.Estimate
    for _, e := range estimates {
        result = append(result, domain.Estimate{
            ID:          e.ID,
            ProductName: e.ProductName,
            Price:       e.Price,
        })
    }

    return result, nil
}

func (r *estimateRepository) GetEstimateByID(id int) (domain.Estimate, error) {
    estimate, err := r.queries.GetEstimateByID(context.Background(), int32(id))
    if err != nil {
        return domain.Estimate{}, err
    }

    return domain.Estimate{
        ID:          estimate.ID,
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    }, nil
}

func (r *estimateRepository) UpdateEstimate(estimate domain.Estimate) (domain.Estimate, error) {
    updatedEstimate, err := r.queries.UpdateEstimate(context.Background(), db.UpdateEstimateParams{
        ID:          int32(estimate.ID),
        ProductName: estimate.ProductName,
        Price:       estimate.Price,
        Longitude:   estimate.Longitude,
        Latitude:    estimate.Latitude,
    })
    if err != nil {
        return domain.Estimate{}, err
    }

    return domain.Estimate{
        ID:          updatedEstimate.ID,
        ProductName: updatedEstimate.ProductName,
        Price:       updatedEstimate.Price,
    }, nil
}

func (r *estimateRepository) DeleteEstimate(id int) error {
    return r.queries.DeleteEstimate(context.Background(), int32(id))
}
