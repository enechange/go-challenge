-- name: GetAllLocations :many
SELECT id, name, address, latitude, longitude, last_updated
FROM locations;

-- name: GetLocationsByDateRange :many
SELECT id, name, address, latitude, longitude, last_updated
FROM locations
WHERE last_updated >= ? AND last_updated < ?;

-- name: GetLocationsByDateFrom :many
SELECT id, name, address, latitude, longitude, last_updated
FROM locations
WHERE last_updated >= ?;

-- name: GetLocationsByDateTo :many
SELECT id, name, address, latitude, longitude, last_updated
FROM locations
WHERE last_updated < ?;

-- name: GetEVSEsByLocationID :many
SELECT id, location_id, uid, status
FROM evses
WHERE location_id = ?;

-- name: GetAllEVSEs :many
SELECT id, location_id, uid, status
FROM evses;
