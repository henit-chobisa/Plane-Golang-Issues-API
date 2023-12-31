// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: projects.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
INSERT INTO "projects" ("id", "name", "description", "workspace_id")
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING id, name, description, workspace_id
`

type CreateProjectParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	WorkspaceID uuid.UUID      `json:"workspace_id"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.Name, arg.Description, arg.WorkspaceID)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.WorkspaceID,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM "projects" WHERE "id" = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, description, workspace_id FROM "projects" WHERE "id" = $1
`

func (q *Queries) GetProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.WorkspaceID,
	)
	return i, err
}

const updateProject = `-- name: UpdateProject :exec
UPDATE "projects"
SET "name" = $2, "description" = $3
WHERE "id" = $1
`

type UpdateProjectParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) error {
	_, err := q.db.ExecContext(ctx, updateProject, arg.ID, arg.Name, arg.Description)
	return err
}
