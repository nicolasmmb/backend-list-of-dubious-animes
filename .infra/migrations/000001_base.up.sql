-- POSTGRESQL
BEGIN;

-- create schema
CREATE SCHEMA IF NOT EXISTS "animes";

-- set schema to animes
SET search_path TO 'animes';

-- create table animes
CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "avatar" VARCHAR(1024) NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);

-- create index
CREATE INDEX "users_name_index" ON "users" ("name");
CREATE INDEX "users_email_index" ON "users" ("email");
CREATE INDEX "users_email_deleted_at_index" ON "users" ("email", "deleted_at");
CREATE INDEX "users_deleted_at_index" ON "users" ("deleted_at");

-- commit transaction
COMMIT;
