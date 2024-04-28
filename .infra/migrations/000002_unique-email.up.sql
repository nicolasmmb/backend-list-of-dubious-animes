-- Desc: ADD unique constraint on email column in users table

SET search_path TO 'animes';
ALTER TABLE users ADD CONSTRAINT users_email_unique UNIQUE (email);