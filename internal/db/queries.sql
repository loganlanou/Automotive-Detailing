-- name: GetAllPackages :many
SELECT * FROM packages
WHERE is_active = 1
ORDER BY sort_order, id;

-- name: GetPackageBySlug :one
SELECT * FROM packages
WHERE slug = ? LIMIT 1;

-- name: ListFeaturedJobs :many
SELECT j.*, v.make, v.model, v.year, v.slug as vehicle_slug
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
WHERE j.featured = 1
ORDER BY j.completed_at DESC
LIMIT ?;

-- name: GetJobByID :one
SELECT * FROM jobs WHERE id = ?;

-- name: ListJobs :many
SELECT j.*, v.make, v.model, v.year, v.slug as vehicle_slug
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
ORDER BY j.completed_at DESC
LIMIT ? OFFSET ?;

-- name: GetVehicleBySlug :one
SELECT * FROM vehicles
WHERE slug = ? LIMIT 1;

-- name: GetMediaForJob :many
SELECT * FROM media
WHERE job_id = ?
ORDER BY sort_order, id;

-- name: GetMediaForVehicle :many
SELECT * FROM media
WHERE vehicle_id = ?
ORDER BY sort_order, id;

-- name: ListReviews :many
SELECT * FROM reviews
ORDER BY created_at DESC
LIMIT ?;

-- name: ListFeaturedReviews :many
SELECT * FROM reviews
WHERE is_featured = 1
ORDER BY created_at DESC
LIMIT ?;

-- name: GetPostBySlug :one
SELECT * FROM posts
WHERE slug = ? AND published_at IS NOT NULL LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
WHERE published_at IS NOT NULL
ORDER BY published_at DESC
LIMIT ? OFFSET ?;
