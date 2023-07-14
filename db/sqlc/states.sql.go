// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: states.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createState = `-- name: CreateState :one
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
) RETURNING created_at, updated_at, id, name, description, color, slug, created_by_id, project_id, workspace_id
`

type CreateStateParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Color       sql.NullString `json:"color"`
	Slug        string         `json:"slug"`
	CreatedByID uuid.UUID      `json:"created_by_id"`
	ProjectID   uuid.UUID      `json:"project_id"`
	WorkspaceID uuid.UUID      `json:"workspace_id"`
}

func (q *Queries) CreateState(ctx context.Context, arg CreateStateParams) (State, error) {
	row := q.db.QueryRowContext(ctx, createState,
		arg.Name,
		arg.Description,
		arg.Color,
		arg.Slug,
		arg.CreatedByID,
		arg.ProjectID,
		arg.WorkspaceID,
	)
	var i State
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Slug,
		&i.CreatedByID,
		&i.ProjectID,
		&i.WorkspaceID,
	)
	return i, err
}

const deleteState = `-- name: DeleteState :exec
DELETE FROM "states" WHERE "id" = $1
`

func (q *Queries) DeleteState(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteState, id)
	return err
}

const getState = `-- name: GetState :one
SELECT created_at, updated_at, id, name, description, color, slug, created_by_id, project_id, workspace_id FROM "states" WHERE "id" = $1
`

func (q *Queries) GetState(ctx context.Context, id uuid.UUID) (State, error) {
	row := q.db.QueryRowContext(ctx, getState, id)
	var i State
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Slug,
		&i.CreatedByID,
		&i.ProjectID,
		&i.WorkspaceID,
	)
	return i, err
}

const updateState = `-- name: UpdateState :exec
UPDATE "states"
SET "name" = $2, "description" = $3, "color" = $4, "slug" = $5
WHERE "id" = $1
`

type UpdateStateParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Color       sql.NullString `json:"color"`
	Slug        string         `json:"slug"`
}

func (q *Queries) UpdateState(ctx context.Context, arg UpdateStateParams) error {
	_, err := q.db.ExecContext(ctx, updateState,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Color,
		arg.Slug,
	)
	return err
}
