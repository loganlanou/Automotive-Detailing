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

-- Admin queries

-- name: CountPackages :one
SELECT COUNT(*) FROM packages;

-- name: CountVehicles :one
SELECT COUNT(*) FROM vehicles;

-- name: CountJobs :one
SELECT COUNT(*) FROM jobs;

-- name: CountMedia :one
SELECT COUNT(*) FROM media;

-- name: CountReviews :one
SELECT COUNT(*) FROM reviews;

-- name: CountPosts :one
SELECT COUNT(*) FROM posts;

-- name: GetAllPackagesAdmin :many
SELECT * FROM packages
ORDER BY sort_order, id;

-- name: GetPackageByID :one
SELECT * FROM packages
WHERE id = ? LIMIT 1;

-- name: CreatePackage :one
INSERT INTO packages (slug, name, short_desc, long_desc, price_min, price_max, duration_est, is_active, sort_order)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdatePackage :one
UPDATE packages
SET slug = ?, name = ?, short_desc = ?, long_desc = ?, price_min = ?, price_max = ?, duration_est = ?, is_active = ?, sort_order = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeletePackage :exec
DELETE FROM packages WHERE id = ?;

-- Work/Portfolio queries

-- name: GetWorkBySlug :one
SELECT
    j.*,
    v.slug as vehicle_slug,
    v.year as vehicle_year,
    v.make as vehicle_make,
    v.model as vehicle_model,
    v.trim as vehicle_trim,
    v.color as vehicle_color,
    v.dealership_name,
    v.dealership_logo_url,
    v.dealership_listing_url,
    v.dealership_location,
    p.name as package_name,
    p.short_desc as package_short_desc,
    p.price_min as package_price_min,
    p.price_max as package_price_max
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
LEFT JOIN packages p ON j.package_id = p.id
WHERE j.slug = ? LIMIT 1;

-- name: GetFeaturedWork :many
SELECT
    j.id,
    j.slug,
    j.featured,
    j.highlight_text,
    j.display_price,
    j.completed_at,
    v.year as vehicle_year,
    v.make as vehicle_make,
    v.model as vehicle_model,
    v.trim as vehicle_trim,
    v.dealership_name,
    p.name as package_name
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
LEFT JOIN packages p ON j.package_id = p.id
WHERE j.featured = 1 AND j.completed_at IS NOT NULL
ORDER BY j.completed_at DESC
LIMIT ?;

-- name: ListAllWork :many
SELECT
    j.id,
    j.slug,
    j.featured,
    j.highlight_text,
    j.display_price,
    j.completed_at,
    v.year as vehicle_year,
    v.make as vehicle_make,
    v.model as vehicle_model,
    v.trim as vehicle_trim,
    v.dealership_name,
    p.name as package_name
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
LEFT JOIN packages p ON j.package_id = p.id
WHERE j.completed_at IS NOT NULL
ORDER BY j.completed_at DESC
LIMIT ? OFFSET ?;

-- name: GetRelatedWork :many
SELECT
    j.id,
    j.slug,
    j.featured,
    j.completed_at,
    v.year as vehicle_year,
    v.make as vehicle_make,
    v.model as vehicle_model,
    v.dealership_name,
    p.name as package_name
FROM jobs j
LEFT JOIN vehicles v ON j.vehicle_id = v.id
LEFT JOIN packages p ON j.package_id = p.id
WHERE j.id != ?
  AND (v.make = ? OR j.package_id = ?)
  AND j.completed_at IS NOT NULL
ORDER BY j.featured DESC, j.completed_at DESC
LIMIT 3;
