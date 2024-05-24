package ports

import "price-estimation-api/internal/domain"

type EstimateRepository interface {
    CreateEstimate(estimate domain.Estimate) (domain.Estimate, error)
    GetEstimates() ([]domain.Estimate, error)
    GetEstimateByID(id int) (domain.Estimate, error)
    UpdateEstimate(estimate domain.Estimate) (domain.Estimate, error)
    DeleteEstimate(id int) error
}
