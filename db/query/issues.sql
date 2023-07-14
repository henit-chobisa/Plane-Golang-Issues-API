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
    json_build_object(
        'name', i.id,
        'description', i.description,
        'priority', i.priority,
        'start_date', i.start_date,
        'target_date', i.target_date,
        'state_detail', json_build_object(
            'name', s.id,
            'color', s.color
        ),
        'user', json_build_object(
            'username', u.username,
            'email', u.email
        ),
        'project', json_build_object(
            'name', p.name,
            'description', p.description
        )
    ) AS issue_data
FROM
    issues i
JOIN
    projects p ON i.project_id = p.id
JOIN
    users u ON i.created_by_id = u.id
JOIN
    states s ON i.state_id = s.id
WHERE
    p.id = $1;

-- name: UpdateIssue :exec
UPDATE "issues"
SET "priority" = $2, "target_date" = $3
WHERE "id" = $1;

-- name: DeleteIssue :exec
DELETE FROM "issues" WHERE "id" = $1;

