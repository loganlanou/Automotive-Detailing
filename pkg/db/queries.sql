-- Package queries

-- name: GetAllPackages :many
SELECT * FROM packages
WHERE is_active = 1
ORDER BY sort_order, id;

-- name: GetPackageBySlug :one
SELECT * FROM packages
WHERE slug = ? LIMIT 1;

-- name: GetPackageByID :one
SELECT * FROM packages
WHERE id = ? LIMIT 1;

-- name: GetAllPackagesAdmin :many
SELECT * FROM packages
ORDER BY sort_order, id;

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

-- name: CountPackages :one
SELECT COUNT(*) FROM packages;

-- Gallery queries

-- name: ListGalleryGroups :many
SELECT * FROM gallery_groups
ORDER BY sort_order, created_at DESC
LIMIT ? OFFSET ?;

-- name: ListFeaturedGalleryGroups :many
SELECT * FROM gallery_groups
WHERE is_featured = 1
ORDER BY sort_order, created_at DESC
LIMIT ?;

-- name: GetGalleryGroupBySlug :one
SELECT * FROM gallery_groups
WHERE slug = ? LIMIT 1;

-- name: GetGalleryGroupByID :one
SELECT * FROM gallery_groups
WHERE id = ? LIMIT 1;

-- name: CreateGalleryGroup :one
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateGalleryGroup :one
UPDATE gallery_groups
SET title = ?, slug = ?, vehicle_make = ?, vehicle_model = ?, vehicle_year = ?, description = ?, is_featured = ?, sort_order = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteGalleryGroup :exec
DELETE FROM gallery_groups WHERE id = ?;

-- name: CountGalleryGroups :one
SELECT COUNT(*) FROM gallery_groups;

-- Media queries

-- name: GetMediaForGalleryGroup :many
SELECT * FROM media
WHERE gallery_group_id = ?
ORDER BY sort_order, id;

-- name: GetHeroImageForGalleryGroup :one
SELECT * FROM media
WHERE gallery_group_id = ? AND kind = 'hero'
ORDER BY sort_order
LIMIT 1;

-- name: CreateMedia :one
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateMedia :one
UPDATE media
SET url = ?, kind = ?, sort_order = ?, alt_text = ?
WHERE id = ?
RETURNING *;

-- name: DeleteMedia :exec
DELETE FROM media WHERE id = ?;

-- name: CountMedia :one
SELECT COUNT(*) FROM media;

-- Review queries

-- name: ListReviews :many
SELECT * FROM reviews
ORDER BY created_at DESC
LIMIT ?;

-- name: ListFeaturedReviews :many
SELECT * FROM reviews
WHERE is_featured = 1
ORDER BY created_at DESC
LIMIT ?;

-- name: CreateReview :one
INSERT INTO reviews (author, rating, body, source, is_featured)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews WHERE id = ?;

-- name: CountReviews :one
SELECT COUNT(*) FROM reviews;

-- Booking queries

-- name: ListBookings :many
SELECT * FROM bookings
ORDER BY requested_start DESC
LIMIT ? OFFSET ?;

-- name: ListBookingsByStatus :many
SELECT * FROM bookings
WHERE status = ?
ORDER BY requested_start ASC
LIMIT ? OFFSET ?;

-- name: ListUpcomingBookings :many
SELECT * FROM bookings
WHERE requested_start >= datetime('now')
  AND status IN ('pending', 'confirmed')
ORDER BY requested_start ASC
LIMIT ?;

-- name: GetBookingByID :one
SELECT * FROM bookings
WHERE id = ? LIMIT 1;

-- name: CreateBooking :one
INSERT INTO bookings (
    customer_name,
    email,
    phone,
    vehicle_details,
    service_interest,
    notes,
    requested_start,
    requested_end,
    status,
    source,
    clerk_user_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateBookingStatus :one
UPDATE bookings
SET status = ?, internal_notes = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: CountBookings :one
SELECT COUNT(*) FROM bookings;

-- name: CountBookingsByStatus :one
SELECT COUNT(*) FROM bookings
WHERE status = ?;

-- name: ListBlockedSlots :many
SELECT requested_start, requested_end, status
FROM bookings
WHERE requested_start >= ?
  AND requested_start < ?
  AND status IN ('pending', 'confirmed')
ORDER BY requested_start;

-- name: CountBlockedSlotsAt :one
SELECT COUNT(*)
FROM bookings
WHERE requested_start = ?
  AND status IN ('pending', 'confirmed');

-- name: ListBookingsForCalendar :many
SELECT id, customer_name, email, phone, vehicle_details, service_interest, requested_start, requested_end, status
FROM bookings
WHERE requested_start >= ?
  AND requested_start < ?
ORDER BY requested_start ASC;
