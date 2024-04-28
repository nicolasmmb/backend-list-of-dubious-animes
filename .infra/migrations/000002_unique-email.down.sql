-- Desc: Remove unique constraint on email column in users table
ALTER TABLE users DROP CONSTRAINT users_email_unique;