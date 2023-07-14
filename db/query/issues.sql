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
SELECT
    i.id AS issue_id,
    i.description AS issue_description,
    i.priority AS issue_priority,
    i.start_date AS issue_start_date,
    i.target_date AS issue_target_date,
    p.id AS project_id,
    p.name AS project_name,
    p.description AS project_description,
    u.id AS created_by_id,
    u.username AS created_by_username,
    u.email AS created_by_email
FROM
    issues i
JOIN
    projects p ON i.project_id = p.id
JOIN
    users u ON i.created_by_id = u.id
WHERE
    p.id = $1;

-- name: UpdateIssue :exec
UPDATE "issues"
SET "priority" = $2, "target_date" = $3
WHERE "id" = $1;

-- name: DeleteIssue :exec
DELETE FROM "issues" WHERE "id" = $1;

