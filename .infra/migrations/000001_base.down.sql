BEGIN;

DROP TABLE IF EXISTS "users";
DROP INDEX IF EXISTS "users_name_index";
DROP INDEX IF EXISTS "users_email_index";
DROP INDEX IF EXISTS "users_email_deleted_at_index";
DROP INDEX IF EXISTS "users_deleted_at_index";

COMMIT;