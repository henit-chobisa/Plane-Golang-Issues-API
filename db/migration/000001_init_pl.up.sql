CREATE TABLE "users" (
    "id" uuid NOT NULL,
    "username" varchar(128) NOT NULL,
    "email" varchar(255),
    PRIMARY KEY ("id")
);


CREATE TABLE "workspaces" (
    "id" uuid NOT NULL,
    "name" varchar(80) NOT NULL,
    "slug" varchar(48) NOT NULL,
    "created_by_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("created_by_id") REFERENCES "users"("id")
);

CREATE TABLE "projects" (
    "id" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "description" text,
    "workspace_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("workspace_id") REFERENCES "workspaces"("id")
);

CREATE TABLE "states" (
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "id" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "description" text,
    "color" varchar(7),
    "slug" varchar(48) NOT NULL,
    "created_by_id" uuid NOT NULL,
    "project_id" uuid NOT NULL,
    "workspace_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("created_by_id") REFERENCES "users"("id"),
    FOREIGN KEY ("project_id") REFERENCES "projects"("id"),
    FOREIGN KEY ("workspace_id") REFERENCES "workspaces"("id")
);

CREATE TABLE "issues" (
    "id" uuid NOT NULL,
    "description" text,
    "priority" int NOT NULL,
    "start_date" date,
    "target_date" date,
    "created_by_id" uuid NOT NULL,
    "project_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("created_by_id") REFERENCES "users"("id"),
    FOREIGN KEY ("project_id") REFERENCES "projects"("id")
);