package domain

type Estimate struct {
    ID           int      `json:"id"`
    ProductName  string   `json:"product_name"`
    Price        float64  `json:"price"`
    Longitude    float64  `json:"longitude"`
    Latitude     float64  `json:"latitude"`
}
