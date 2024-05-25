package domain

type Estimate struct {
    ID           int      `json:"id"`
    ProductName  string   `json:"product_name"`
    Price        float64  `json:"price"`
    Longitude    string  `json:"longitude"`
    Latitude     string  `json:"latitude"`
}
