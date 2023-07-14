
-- name: CreateWorkspace :one
INSERT INTO "workspaces" ("id", "name", "slug", "created_by_id")
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING *;

-- name: GetWorkspace :one
SELECT * FROM "workspaces" WHERE "id" = $1;

-- name: UpdateWorkspace :exec
UPDATE "workspaces"
SET "name" = $2, "slug" = $3
WHERE "id" = $1;

