
-- name: CreateProject :one
INSERT INTO "projects" ("id", "name", "description", "workspace_id")
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING *;

-- name: GetProject :one 
SELECT * FROM "projects" WHERE "id" = $1;

-- name: UpdateProject :exec
UPDATE "projects"
SET "name" = $2, "description" = $3
WHERE "id" = $1;

-- name: DeleteProject :exec
DELETE FROM "projects" WHERE "id" = $1;