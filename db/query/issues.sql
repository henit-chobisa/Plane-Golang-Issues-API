-- name: CreateIssue :one
INSERT INTO "issues" (
    "id",
    "description",
    "priority",
    "start_date",
    "target_date",
    "created_by_id",
    "project_id"
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetIssue :one
SELECT * FROM "issues"
WHERE id=$1 LIMIT 1;

-- name: ListIssuesByProject :many
SELECT * FROM "issues"
WHERE project_id=$1
ORDER BY priority;

-- name: UpdateIssue :exec
UPDATE "issues"
SET "priority" = $2, "target_date" = $3
WHERE "id" = $1;

-- name: DeleteIssue :exec
DELETE FROM "issues" WHERE "id" = $1;

