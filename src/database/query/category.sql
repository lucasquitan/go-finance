-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    title,
    type,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: getCategory :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: getCategories :many
SELECT * FROM categories WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4;

-- name: updateCategory :one
UPDATE categories SET title = $2, description = $3 WHERE id = $1 RETURNING *;

-- name: deleteCategory :exec
DELETE FROM categories WHERE id = $1;