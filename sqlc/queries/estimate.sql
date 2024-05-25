-- name: CreateEstimate :one
INSERT INTO estimates (product_name, price, longitude, latitude)
VALUES ($1, $2, $3, $4)
-- RETURNING id, product_name, price, longitude, latitude;

-- name: GetEstimatesWithPagination :many
SELECT id, product_name, price, longitude, latitude
FROM estimates
ORDER BY id
-- LIMIT $1 OFFSET $2;

-- name: GetEstimateByID :one
SELECT id, product_name, price, longitude, latitude
FROM estimates
WHERE id = $1;

-- name: UpdateEstimate :one
UPDATE estimates
SET product_name = $2, price = $3, longitude = $4, latitude = $5
WHERE id = $1
-- RETURNING id, product_name, price, longitude, latitude;

-- name: DeleteEstimate :exec
DELETE FROM estimates
WHERE id = $1;
