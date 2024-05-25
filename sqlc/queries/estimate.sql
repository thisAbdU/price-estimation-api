-- name: CreateEstimate :one
INSERT INTO estimates (product_name, price, longitude, latitude)
VALUES ($1, $2, $3, $4)

-- name: DeleteEstimate :exec
DELETE FROM estimates
WHERE id = $1;

-- name: GetEstimateByID :one
SELECT id, product_name, price, longitude, latitude
FROM estimates
WHERE id = $1;

-- name: GetEstimatesWithPagination :many
SELECT id, product_name, price, longitude, latitude
FROM estimates
ORDER BY id

-- name: UpdateEstimate :one
UPDATE estimates
SET product_name = $2, price = $3, longitude = $4, latitude = $5
WHERE id = $1;




