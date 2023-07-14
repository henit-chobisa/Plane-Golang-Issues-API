-- name: CreateState :one
INSERT INTO "states" (
    "created_at",
    "updated_at",
    "id",
    "name",
    "description",
    "color",
    "slug",
    "created_by_id",
    "project_id",
    "workspace_id"
)
VALUES (
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;

-- name: GetState :one
SELECT * FROM "states" WHERE "id" = $1;

-- name: UpdateState :exec
UPDATE "states"
SET "name" = $2, "description" = $3, "color" = $4, "slug" = $5
WHERE "id" = $1;

-- name: DeleteState :exec
DELETE FROM "states" WHERE "id" = $1;